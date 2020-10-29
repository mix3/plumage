package example_test

import (
	"fmt"
	"testing"
	"time"

	"example"
	mock_example "example/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func pstr(v string) *string {
	return &v
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

type c1 struct{}

func (c c1) ID() example.ID {
	return example.ID("a")
}

func (c c1) Stringer() fmt.Stringer {
	return stringer("stringer")
}

func (c c1) Hoge() time.Time {
	return time.Time{}
}

func (c c1) Foo() {
	panic(fmt.Errorf("not implemented"))
}

func (c c1) Bar(_ string) {
	panic(fmt.Errorf("not implemented"))
}

type c2 struct{}

func (c c2) ID() example.ID {
	return example.ID("b")
}

func (c c2) Fuga() *string {
	return pstr("fuga")
}

func (c c2) Piyo() *time.Time {
	return &time.Time{}
}

func (c c2) Baz() (string, string) {
	panic(fmt.Errorf("not implemented"))
}

func TestExample(t *testing.T) {
	ctrl := gomock.NewController(t)

	child1 := mock_example.NewMockChild1(ctrl)
	child1.EXPECT().ID().AnyTimes().Return(example.ID("a"))
	child1.EXPECT().Stringer().AnyTimes().Return(stringer("stringer"))
	child1.EXPECT().Hoge().AnyTimes().Return(time.Time{})

	child2 := mock_example.NewMockChild2(ctrl)
	child2.EXPECT().ID().AnyTimes().Return(example.ID("b"))
	child2.EXPECT().Fuga().AnyTimes().Return(pstr("fuga"))
	child2.EXPECT().Piyo().AnyTimes().Return(&time.Time{})

	assert.Empty(t, cmp.Diff(example.NewBaseStruct(child1), example.NewBaseStruct(c1{})))
	assert.Empty(t, cmp.Diff(example.NewBaseStruct(child2), example.NewBaseStruct(c2{})))
	assert.NotEmpty(t, cmp.Diff(example.NewBaseStruct(child1), example.NewBaseStruct(c2{})))
	assert.NotEmpty(t, cmp.Diff(example.NewBaseStruct(child2), example.NewBaseStruct(c1{})))
}
