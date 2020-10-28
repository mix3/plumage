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
}

type Child2 interface {
	Base
	Fuga() []string
	Piyo() *time.Time
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
					Name:        "ID",
					PkgName:     "",
					TypeName:    "ID",
					IsArray:     false,
					ArrayCount:  0,
					IsStar:      false,
					IsObj:       true,
					IsInterface: false,
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
					Name:        "ID",
					PkgName:     "",
					TypeName:    "ID",
					IsArray:     false,
					ArrayCount:  0,
					IsStar:      false,
					IsObj:       true,
					IsInterface: false,
				},
				{
					Name:        "Hoge",
					PkgName:     "",
					TypeName:    "int",
					IsArray:     false,
					ArrayCount:  0,
					IsStar:      false,
					IsObj:       false,
					IsInterface: false,
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
					Name:        "ID",
					PkgName:     "",
					TypeName:    "ID",
					IsArray:     false,
					ArrayCount:  0,
					IsStar:      false,
					IsObj:       true,
					IsInterface: false,
				},
				{
					Name:        "Fuga",
					PkgName:     "",
					TypeName:    "string",
					IsArray:     true,
					ArrayCount:  1,
					IsStar:      false,
					IsObj:       false,
					IsInterface: false,
				},
				{
					Name:        "Piyo",
					PkgName:     "time",
					TypeName:    "Time",
					IsArray:     false,
					ArrayCount:  0,
					IsStar:      true,
					IsObj:       false,
					IsInterface: false,
				},
			},
		},
	}
	assert.Empty(t, cmp.Diff(want, got))
}
