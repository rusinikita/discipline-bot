package dbtest

import (
	"github.com/stretchr/testify/mock"
	"reflect"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) OnList() *mock.Call {
	return m.On("List")
}

func (m *MockDB) List(table string, list interface{}, view string) error {
	ret := m.Called()

	reflect.ValueOf(list).Elem().Set(reflect.ValueOf(ret.Get(0)))

	return nil
}
