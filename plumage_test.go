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
}
`)
	if err := ioutil.WriteFile(filepath.Join(dir, "example.go"), content, 0644); err != nil {
		t.Error(err)
		return
	}

	got, err := inspect(dir)
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
			Child:       []string{},
			FieldInfos:  []*FieldInfo{},
		},
		&TypeInfo{
			PkgName:     "example",
			TypeName:    "",
			Name:        "Base",
			IsInterface: true,
			IsArray:     false,
			ArrayCount:  0,
			Parent:      []string{},
			Child:       []string{"Child1", "Child2"},
			FieldInfos: []*FieldInfo{
				{
					Name:   "ID",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
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
			Child:       []string{},
			FieldInfos: []*FieldInfo{
				{
					Name:   "ID",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
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
			Child:       []string{},
			FieldInfos: []*FieldInfo{
				{
					Name:   "ID",
					Params: []*Field{},
					Results: []*Field{
						{
							PkgName:    "",
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
			},
		},
	}
	assert.Empty(t, cmp.Diff(want, got))
}
