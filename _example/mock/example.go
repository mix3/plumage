// Code generated by MockGen. DO NOT EDIT.
// Source: example.go

// Package mock_example is a generated GoMock package.
package mock_example

import (
	example "example"
	fmt "fmt"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockBase is a mock of Base interface
type MockBase struct {
	ctrl     *gomock.Controller
	recorder *MockBaseMockRecorder
}

// MockBaseMockRecorder is the mock recorder for MockBase
type MockBaseMockRecorder struct {
	mock *MockBase
}

// NewMockBase creates a new mock instance
func NewMockBase(ctrl *gomock.Controller) *MockBase {
	mock := &MockBase{ctrl: ctrl}
	mock.recorder = &MockBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBase) EXPECT() *MockBaseMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockBase) ID() example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(example.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockBaseMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockBase)(nil).ID))
}

// MockChild1 is a mock of Child1 interface
type MockChild1 struct {
	ctrl     *gomock.Controller
	recorder *MockChild1MockRecorder
}

// MockChild1MockRecorder is the mock recorder for MockChild1
type MockChild1MockRecorder struct {
	mock *MockChild1
}

// NewMockChild1 creates a new mock instance
func NewMockChild1(ctrl *gomock.Controller) *MockChild1 {
	mock := &MockChild1{ctrl: ctrl}
	mock.recorder = &MockChild1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChild1) EXPECT() *MockChild1MockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockChild1) ID() example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(example.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockChild1MockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockChild1)(nil).ID))
}

// Stringer mocks base method
func (m *MockChild1) Stringer() fmt.Stringer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stringer")
	ret0, _ := ret[0].(fmt.Stringer)
	return ret0
}

// Stringer indicates an expected call of Stringer
func (mr *MockChild1MockRecorder) Stringer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stringer", reflect.TypeOf((*MockChild1)(nil).Stringer))
}

// Hoge mocks base method
func (m *MockChild1) Hoge() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hoge")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Hoge indicates an expected call of Hoge
func (mr *MockChild1MockRecorder) Hoge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hoge", reflect.TypeOf((*MockChild1)(nil).Hoge))
}

// Foo mocks base method
func (m *MockChild1) Foo() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo")
}

// Foo indicates an expected call of Foo
func (mr *MockChild1MockRecorder) Foo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockChild1)(nil).Foo))
}

// Bar mocks base method
func (m *MockChild1) Bar(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Bar", arg0)
}

// Bar indicates an expected call of Bar
func (mr *MockChild1MockRecorder) Bar(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bar", reflect.TypeOf((*MockChild1)(nil).Bar), arg0)
}

// MockChild2 is a mock of Child2 interface
type MockChild2 struct {
	ctrl     *gomock.Controller
	recorder *MockChild2MockRecorder
}

// MockChild2MockRecorder is the mock recorder for MockChild2
type MockChild2MockRecorder struct {
	mock *MockChild2
}

// NewMockChild2 creates a new mock instance
func NewMockChild2(ctrl *gomock.Controller) *MockChild2 {
	mock := &MockChild2{ctrl: ctrl}
	mock.recorder = &MockChild2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChild2) EXPECT() *MockChild2MockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockChild2) ID() example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(example.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockChild2MockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockChild2)(nil).ID))
}

// Fuga mocks base method
func (m *MockChild2) Fuga() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fuga")
	ret0, _ := ret[0].(*string)
	return ret0
}

// Fuga indicates an expected call of Fuga
func (mr *MockChild2MockRecorder) Fuga() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fuga", reflect.TypeOf((*MockChild2)(nil).Fuga))
}

// Piyo mocks base method
func (m *MockChild2) Piyo() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Piyo")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// Piyo indicates an expected call of Piyo
func (mr *MockChild2MockRecorder) Piyo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Piyo", reflect.TypeOf((*MockChild2)(nil).Piyo))
}

// Baz mocks base method
func (m *MockChild2) Baz() (string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Baz")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// Baz indicates an expected call of Baz
func (mr *MockChild2MockRecorder) Baz() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Baz", reflect.TypeOf((*MockChild2)(nil).Baz))
}

// MockX is a mock of X interface
type MockX struct {
	ctrl     *gomock.Controller
	recorder *MockXMockRecorder
}

// MockXMockRecorder is the mock recorder for MockX
type MockXMockRecorder struct {
	mock *MockX
}

// NewMockX creates a new mock instance
func NewMockX(ctrl *gomock.Controller) *MockX {
	mock := &MockX{ctrl: ctrl}
	mock.recorder = &MockXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockX) EXPECT() *MockXMockRecorder {
	return m.recorder
}

// Int mocks base method
func (m *MockX) Int() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int")
	ret0, _ := ret[0].(int)
	return ret0
}

// Int indicates an expected call of Int
func (mr *MockXMockRecorder) Int() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockX)(nil).Int))
}

// StarInt mocks base method
func (m *MockX) StarInt() *int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarInt")
	ret0, _ := ret[0].(*int)
	return ret0
}

// StarInt indicates an expected call of StarInt
func (mr *MockXMockRecorder) StarInt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarInt", reflect.TypeOf((*MockX)(nil).StarInt))
}

// IntList mocks base method
func (m *MockX) IntList() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntList")
	ret0, _ := ret[0].([]int)
	return ret0
}

// IntList indicates an expected call of IntList
func (mr *MockXMockRecorder) IntList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntList", reflect.TypeOf((*MockX)(nil).IntList))
}

// StarIntList mocks base method
func (m *MockX) StarIntList() []*int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarIntList")
	ret0, _ := ret[0].([]*int)
	return ret0
}

// StarIntList indicates an expected call of StarIntList
func (mr *MockXMockRecorder) StarIntList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarIntList", reflect.TypeOf((*MockX)(nil).StarIntList))
}

// IntListList mocks base method
func (m *MockX) IntListList() [][]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntListList")
	ret0, _ := ret[0].([][]int)
	return ret0
}

// IntListList indicates an expected call of IntListList
func (mr *MockXMockRecorder) IntListList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntListList", reflect.TypeOf((*MockX)(nil).IntListList))
}

// StarIntListList mocks base method
func (m *MockX) StarIntListList() [][]*int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarIntListList")
	ret0, _ := ret[0].([][]*int)
	return ret0
}

// StarIntListList indicates an expected call of StarIntListList
func (mr *MockXMockRecorder) StarIntListList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarIntListList", reflect.TypeOf((*MockX)(nil).StarIntListList))
}

// Time mocks base method
func (m *MockX) Time() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Time indicates an expected call of Time
func (mr *MockXMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockX)(nil).Time))
}

// StarTime mocks base method
func (m *MockX) StarTime() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarTime")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// StarTime indicates an expected call of StarTime
func (mr *MockXMockRecorder) StarTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarTime", reflect.TypeOf((*MockX)(nil).StarTime))
}

// TimeList mocks base method
func (m *MockX) TimeList() []time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeList")
	ret0, _ := ret[0].([]time.Time)
	return ret0
}

// TimeList indicates an expected call of TimeList
func (mr *MockXMockRecorder) TimeList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeList", reflect.TypeOf((*MockX)(nil).TimeList))
}

// StarTimeList mocks base method
func (m *MockX) StarTimeList() []*time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarTimeList")
	ret0, _ := ret[0].([]*time.Time)
	return ret0
}

// StarTimeList indicates an expected call of StarTimeList
func (mr *MockXMockRecorder) StarTimeList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarTimeList", reflect.TypeOf((*MockX)(nil).StarTimeList))
}

// TimeListList mocks base method
func (m *MockX) TimeListList() [][]time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeListList")
	ret0, _ := ret[0].([][]time.Time)
	return ret0
}

// TimeListList indicates an expected call of TimeListList
func (mr *MockXMockRecorder) TimeListList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeListList", reflect.TypeOf((*MockX)(nil).TimeListList))
}

// StarTimeListList mocks base method
func (m *MockX) StarTimeListList() [][]*time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarTimeListList")
	ret0, _ := ret[0].([][]*time.Time)
	return ret0
}

// StarTimeListList indicates an expected call of StarTimeListList
func (mr *MockXMockRecorder) StarTimeListList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarTimeListList", reflect.TypeOf((*MockX)(nil).StarTimeListList))
}

// ID mocks base method
func (m *MockX) ID() example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(example.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockXMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockX)(nil).ID))
}

// StarID mocks base method
func (m *MockX) StarID() *example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarID")
	ret0, _ := ret[0].(*example.ID)
	return ret0
}

// StarID indicates an expected call of StarID
func (mr *MockXMockRecorder) StarID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarID", reflect.TypeOf((*MockX)(nil).StarID))
}

// IDList mocks base method
func (m *MockX) IDList() []example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDList")
	ret0, _ := ret[0].([]example.ID)
	return ret0
}

// IDList indicates an expected call of IDList
func (mr *MockXMockRecorder) IDList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDList", reflect.TypeOf((*MockX)(nil).IDList))
}

// StarIDList mocks base method
func (m *MockX) StarIDList() []*example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarIDList")
	ret0, _ := ret[0].([]*example.ID)
	return ret0
}

// StarIDList indicates an expected call of StarIDList
func (mr *MockXMockRecorder) StarIDList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarIDList", reflect.TypeOf((*MockX)(nil).StarIDList))
}

// IDListList mocks base method
func (m *MockX) IDListList() [][]example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDListList")
	ret0, _ := ret[0].([][]example.ID)
	return ret0
}

// IDListList indicates an expected call of IDListList
func (mr *MockXMockRecorder) IDListList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDListList", reflect.TypeOf((*MockX)(nil).IDListList))
}

// StarIDListList mocks base method
func (m *MockX) StarIDListList() [][]*example.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StarIDListList")
	ret0, _ := ret[0].([][]*example.ID)
	return ret0
}

// StarIDListList indicates an expected call of StarIDListList
func (mr *MockXMockRecorder) StarIDListList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StarIDListList", reflect.TypeOf((*MockX)(nil).StarIDListList))
}

// Base mocks base method
func (m *MockX) Base() example.Base {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Base")
	ret0, _ := ret[0].(example.Base)
	return ret0
}

// Base indicates an expected call of Base
func (mr *MockXMockRecorder) Base() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Base", reflect.TypeOf((*MockX)(nil).Base))
}

// BaseList mocks base method
func (m *MockX) BaseList() []example.Base {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseList")
	ret0, _ := ret[0].([]example.Base)
	return ret0
}

// BaseList indicates an expected call of BaseList
func (mr *MockXMockRecorder) BaseList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseList", reflect.TypeOf((*MockX)(nil).BaseList))
}
