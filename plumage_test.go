package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestPlumage(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	content := []byte(`package example

type ID string

type A interface {
	B() int
}

type Base interface {
	ID() ID
}

type Child1 interface {
	Base
	Hoge() int
	Foo()
	Bar() (string, []int)
}

type Child2 interface {
	Base
	Fuga() []string
	Piyo() *time.Time
	Baz(*time.Time)
	A() A
}
`)
	if err := ioutil.WriteFile(filepath.Join(dir, "example.go"), content, 0644); err != nil {
		t.Error(err)
		return
	}

	got, err := inspect(dir, false)
	if err != nil {
		t.Error(err)
		return
	}

	want := TypeInfos{
		&TypeInfo{
			PkgName:     "example",
			TypeName:    "string",
			Name:        "ID",
			IsInterface: false,
			IsArray:     false,
			ArrayCount:  0,
			Parent:      []string{},
			Child:       []*Name{},
			FieldInfos:  []*FieldInfo{},
		},
		&TypeInfo{
			PkgName:     "example",
			TypeName:    "",
			Name:        "A",
			IsInterface: true,
			IsArray:     false,
			ArrayCount:  0,
			Parent:      []string{},
			Child:       []*Name{},
			FieldInfos: []*FieldInfo{
				{
					Name:   "B",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
							TypeName:   "int",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      false,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
			},
		},
		&TypeInfo{
			PkgName:     "example",
			TypeName:    "",
			Name:        "Base",
			IsInterface: true,
			IsArray:     false,
			ArrayCount:  0,
			Parent:      []string{},
			Child: []*Name{
				{PkgName: "example", Name: "Child1"},
				{PkgName: "example", Name: "Child2"},
			},
			FieldInfos: []*FieldInfo{
				{
					Name:   "ID",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "example",
							TypeName:   "ID",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      true,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
			},
		},
		&TypeInfo{
			PkgName:     "example",
			TypeName:    "",
			Name:        "Child1",
			IsInterface: true,
			IsArray:     false,
			ArrayCount:  0,
			Parent:      []string{"Base"},
			Child:       []*Name{},
			FieldInfos: []*FieldInfo{
				{
					Name:   "ID",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "example",
							TypeName:   "ID",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      true,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
				{
					Name:   "Hoge",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
							TypeName:   "int",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      false,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
				{
					Name:        "Foo",
					Params:      []*Field{},
					Results:     []*Field{},
					IsInterface: false,
					IsValue:     false,
				},
				{
					Name:   "Bar",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
							TypeName:   "string",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      false,
						},
						{
							PkgName:    "",
							TypeName:   "int",
							IsArray:    true,
							ArrayCount: 1,
							IsStar:     false,
							IsObj:      false,
						},
					},
					IsInterface: false,
					IsValue:     false,
				},
			},
		},
		&TypeInfo{
			PkgName:     "example",
			TypeName:    "",
			Name:        "Child2",
			IsInterface: true,
			IsArray:     false,
			ArrayCount:  0,
			Parent:      []string{"Base"},
			Child:       []*Name{},
			FieldInfos: []*FieldInfo{
				{
					Name:   "ID",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "example",
							TypeName:   "ID",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      true,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
				{
					Name:   "Fuga",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
							TypeName:   "string",
							IsArray:    true,
							ArrayCount: 1,
							IsStar:     false,
							IsObj:      false,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
				{
					Name:   "Piyo",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "time",
							TypeName:   "Time",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     true,
							IsObj:      false,
						},
					},
					IsInterface: false,
					IsValue:     true,
				},
				{
					Name: "Baz",
					Params: []*Field{
						{
							PkgName:    "time",
							TypeName:   "Time",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     true,
							IsObj:      false,
						},
					},
					Results:     []*Field{},
					IsInterface: false,
					IsValue:     false,
				},
				{
					Name:   "A",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "example",
							TypeName:   "A",
							IsArray:    false,
							ArrayCount: 0,
							IsStar:     false,
							IsObj:      true,
						},
					},
					IsInterface: true,
					IsValue:     true,
				},
			},
		},
	}
	assert.Empty(t, cmp.Diff(want, got))

	got, err = inspect(dir, true)
	if err != nil {
		t.Error(err)
		return
	}
	want[2].Child = []*Name{{Name: "Child1"}, {Name: "Child2"}}
	want[2].FieldInfos[0].Results[0].PkgName = ""
	want[3].FieldInfos[0].Results[0].PkgName = ""
	want[4].FieldInfos[0].Results[0].PkgName = ""
	want[4].FieldInfos[4].Results[0].PkgName = ""
	assert.Empty(t, cmp.Diff(want, got))
}
