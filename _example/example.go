//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE
//go:generate go run github.com/mix3/plumage

package example

import (
	"fmt"
	"time"
)

type ID string

type IDs []ID

type Base interface {
	ID() ID
}

type Child1 interface {
	Base
	Stringer() fmt.Stringer
	Hoge() time.Time
}

type Child2 interface {
	Base
	Fuga() *string
	Piyo() *time.Time
}

type X interface {
	Int() int
	StarInt() *int
	IntList() []int
	StarIntList() []*int
	IntListList() [][]int
	StarIntListList() [][]*int
	Time() time.Time
	StarTime() *time.Time
	TimeList() []time.Time
	StarTimeList() []*time.Time
	TimeListList() [][]time.Time
	StarTimeListList() [][]*time.Time
	ID() ID
	StarID() *ID
	IDList() []ID
	StarIDList() []*ID
	IDListList() [][]ID
	StarIDListList() [][]*ID
	Base() Base
	BaseList() []Base
}
