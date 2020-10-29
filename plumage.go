package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

var tmpl = template.Must(template.New("tmpl").Funcs(
	template.FuncMap{},
).Parse(`// Code generated by plumage - DO NOT EDIT.

package {{ .PkgName }}

{{- range .TypeInfos.Interfaces }}{{ $TypeInfo := . }}
type {{ .Name }} struct {
	{{- range .FieldInfos.ValueFieldInfos }}
		{{ .Name }}_ {{ .ResultsStr }}
	{{- end }}
}

{{- range .FieldInfos }}
func (v {{ $TypeInfo.Name }}) {{ .Name }}({{ .ParamsStr }}) ({{ .ResultsStr }}) {
	{{- if .IsValue }}
	return v.{{ .Name }}_
	{{- else }}
	panic(fmt.Errorf("Not supported."))
	{{- end }}
}

{{ end }}{{ end }}

{{- range .TypeInfos.Interfaces }}
func New{{ .Name }}(v {{ .FullName }}) {{ .FullName }} {
	if v == nil {
		return nil
	}
	switch {{ if ne 0 .ChildCount }}t := {{ end }}v.(type) {
	{{- range .Child }}
	case {{ .FullName }}:
		return New{{ .Name }}(t)
	{{- end }}
	default:
		return {{ .Name }}{
			{{- range .FieldInfos.ValueFieldInfos }}
				{{- if .IsInterface }}{{ $Field := index .Results 0 }}
					{{- if lt 0 $Field.ArrayCount }}
						{{ .Name }}_: New{{ $Field.TypeName }}List(v.{{ .Name }}()),
					{{- else }}
						{{ .Name }}_: New{{ $Field.TypeName }}(v.{{ .Name }}()),
					{{- end }}
				{{- else }}
					{{ .Name }}_: v.{{ .Name }}(),
				{{- end }}
			{{- end }}
		}
	}
}

func New{{ .Name }}List(vs []{{ .FullName }}) []{{ .FullName }} {
	ret := make([]{{ .FullName }}, 0, len(vs))
	for _, v := range vs {
		ret = append(ret, New{{ .Name }}(v))
	}
	return ret
}

{{ end }}
`))

func main() {
	var src, dir, suffix string
	flag.StringVar(&src, "src", "", "code generate source")
	flag.StringVar(&dir, "dst", "", "output directory")
	flag.StringVar(&suffix, "suffix", ".gen.go", "output file suffix. default .gen.go")
	flag.Parse()

	if src == "" {
		src = filepath.Dir(os.Getenv("GOFILE"))
	}

	if err := run(src, dir, suffix); err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
}

func run(src, dir, suffix string) error {
	tinfos, err := inspect(src, dir == "")
	if err != nil {
		return err
	}

	var fileName, pkgName string
	for _, tinfo := range tinfos {
		pkgName = tinfo.PkgName
		fileName = pkgName
		break
	}
	if dir != "" {
		pkgName = filepath.Base(dir)
	}

	buf := &bytes.Buffer{}
	tmpl.Execute(buf, struct {
		TypeInfos TypeInfos
		PkgName   string
	}{
		TypeInfos: tinfos,
		PkgName:   pkgName,
	})

	// format と import を goimport に丸投げ
	formattedBytes, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return fmt.Errorf("RUn: fail to goimport: source=\n%s\n, error=%w", buf.String(), err)
	}
	formatted := bytes.NewBuffer(formattedBytes)

	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("Run: fail to create directory: path=%s, error=%w", dir, err)
		}
	}

	dst := filepath.Join(src, dir, fileName+suffix)
	f, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Run: fail to create file: filename=%s, error=%w", dst, err)
	}
	defer f.Close()
	_, err = formatted.WriteTo(f)
	if err != nil {
		return fmt.Errorf("Run: fail to write file: filename=%s, error=%w", dst, err)
	}

	return nil
}

func inspect(src string, samePkg bool) (TypeInfos, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, src, func(f os.FileInfo) bool {
		return strings.HasSuffix(f.Name(), ".go") && !strings.HasSuffix(f.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("Run: fail to parse from dir: src=%s, error=%w", src, err)
	}

	tinfos := TypeInfos{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				gd, ok := decl.(*ast.GenDecl)
				if !ok || gd.Tok != token.TYPE {
					continue
				}
				for _, spec := range gd.Specs {
					tspec := spec.(*ast.TypeSpec)
					tinfo := typ(tspec.Type, NewTypeInfo(pkg.Name, tspec.Name.Name))
					tinfos = append(tinfos, tinfo)
				}
			}
		}
	}

	tinfos = tinfos.Post(samePkg)

	return tinfos, nil
}

type TypeInfos []*TypeInfo

func (ts TypeInfos) Interfaces() []*TypeInfo {
	ret := make([]*TypeInfo, 0, len(ts))
	for _, t := range ts {
		if t.IsInterface {
			ret = append(ret, t)
		}
	}
	return ret
}

func (ts TypeInfos) Post(samePkg bool) TypeInfos {
	imap := make(map[string]struct{}, len(ts))
	tmap := make(map[string]struct{}, len(ts))
	for _, t := range ts {
		tmap[t.Name] = struct{}{}
		if t.IsInterface {
			imap[t.Name] = struct{}{}
		}
	}

	p2cs := map[string][]*Name{}
	for _, t := range ts.Interfaces() {
		for _, p := range t.Parent {
			n := &Name{Name: t.Name}
			if !samePkg {
				n.PkgName = t.PkgName
			}
			p2cs[p] = append(p2cs[p], n)
		}
	}

	for _, iface := range ts.Interfaces() {
		// フィールドが interface を返すのか struct を返すのかを区別したい
		for _, finfo := range iface.FieldInfos {
			if !finfo.IsValue {
				continue
			}
			f := finfo.Results[0]
			if _, ok := tmap[f.TypeName]; ok && !samePkg {
				f.PkgName = iface.PkgName
			}
			if _, ok := imap[f.TypeName]; ok {
				finfo.IsInterface = true
				if 2 <= f.ArrayCount {
					panic(fmt.Errorf("Multi-dimensional arrays are not supported"))
				}
			}
		}
		// Child埋める
		if cs, ok := p2cs[iface.Name]; ok {
			iface.Child = cs
		}
	}

	return ts
}

type TypeInfo struct {
	PkgName     string
	TypeName    string
	Name        string
	IsInterface bool
	IsArray     bool
	ArrayCount  int
	Parent      []string
	Child       []*Name
	FieldInfos  FieldInfos
}

type Name struct {
	PkgName string
	Name    string
}

func (n Name) FullName() string {
	if n.PkgName != "" {
		return n.PkgName + "." + n.Name
	}
	return n.Name
}

func (t TypeInfo) ChildCount() int {
	return len(t.Child)
}

func (t TypeInfo) FullName() string {
	if t.PkgName != "" {
		return t.PkgName + "." + t.Name
	}
	return t.Name
}

func NewTypeInfo(pkgName, name string) *TypeInfo {
	return &TypeInfo{
		PkgName:    pkgName,
		Name:       name,
		Parent:     make([]string, 0, 10),
		Child:      make([]*Name, 0, 10),
		FieldInfos: make(FieldInfos, 0, 10),
	}
}

func typ(expr ast.Expr, info *TypeInfo) *TypeInfo {
	switch t := expr.(type) {
	case *ast.Ident:
		info.TypeName = t.Name
	case *ast.ArrayType:
		info.IsArray = true
		info.ArrayCount++
		info = typ(t.Elt, info)
	case *ast.InterfaceType:
		info.IsInterface = true
		for _, method := range t.Methods.List {
			switch mt := method.Type.(type) {
			case *ast.Ident:
				// interface in interface
				tspec, ok := mt.Obj.Decl.(*ast.TypeSpec)
				if !ok {
					return info
				}
				info.Parent = append(info.Parent, mt.Name)
				info = typ(tspec.Type, info)
			case *ast.FuncType:
				// 引数取らず値を一つだけ返すやつに絞る
				mname := method.Names[0].Name
				finfo := NewFieldInfo(mname)
				if 0 < mt.Params.NumFields() {
					for _, p := range mt.Params.List {
						finfo.Params = append(finfo.Params, field(p.Type, &Field{}))
					}
				}
				if 0 < mt.Results.NumFields() {
					for _, r := range mt.Results.List {
						finfo.Results = append(finfo.Results, field(r.Type, &Field{}))
					}
				}
				finfo.IsValue = mt.Params.NumFields() == 0 && mt.Results.NumFields() == 1
				info.FieldInfos = append(info.FieldInfos, finfo)
			default:
				panic(fmt.Errorf("method.Type(%T) does not support", mt))
			}
		}
	}
	return info
}

type FieldInfos []*FieldInfo

func (fs FieldInfos) ValueFieldInfos() FieldInfos {
	ret := make(FieldInfos, 0, len(fs))
	for _, f := range fs {
		if f.IsValue {
			ret = append(ret, f)
		}
	}
	return ret
}

type FieldInfo struct {
	Name        string
	Params      []*Field
	Results     []*Field
	IsValue     bool
	IsInterface bool
}

func NewFieldInfo(name string) *FieldInfo {
	return &FieldInfo{
		Name:    name,
		Params:  make([]*Field, 0, 10),
		Results: make([]*Field, 0, 10),
	}
}

func (f *FieldInfo) ParamsStr() string {
	ret := make([]string, 0, len(f.Params))
	for _, p := range f.Params {
		ret = append(ret, p.FullTypeName())
	}
	return strings.Join(ret, ", ")
}

func (f *FieldInfo) ResultsStr() string {
	ret := make([]string, 0, len(f.Results))
	for _, f := range f.Results {
		ret = append(ret, f.FullTypeName())
	}
	return strings.Join(ret, ", ")
}

type Field struct {
	PkgName    string
	TypeName   string
	IsArray    bool
	ArrayCount int
	IsStar     bool
	IsObj      bool
}

func (f Field) FullTypeName() string {
	str := ""
	if f.IsArray {
		str += strings.Repeat("[]", f.ArrayCount)
	}
	if f.IsStar {
		str += "*"
	}
	if f.PkgName != "" {
		str += f.PkgName + "."
	}
	str += f.TypeName
	return str
}

func field(expr ast.Expr, f *Field) *Field {
	switch rt := expr.(type) {
	case *ast.Ident:
		f.TypeName = rt.Name
		f.IsObj = rt.Obj != nil
	case *ast.ArrayType:
		f.IsArray = true
		f.ArrayCount++
		f = field(rt.Elt, f)
	case *ast.StarExpr:
		f.IsStar = true
		f = field(rt.X, f)
	case *ast.SelectorExpr:
		x, ok := rt.X.(*ast.Ident)
		if !ok {
			break
		}
		f.PkgName = x.Name
		f = field(rt.Sel, f)
	}
	return f
}
