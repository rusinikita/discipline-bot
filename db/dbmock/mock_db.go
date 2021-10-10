package dbmock

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) OnOne() *mock.Call {
	return m.On("One")
}

func (m *MockDB) OnList() *mock.Call {
	return m.On("List")
}

func (m *MockDB) OnPatch() *mock.Call {
	return m.On("Patch")
}

func (m *MockDB) One(table, id string, entity interface{}) error {
	return m.Called().Get(0).(error)
}

func (m *MockDB) List(table string, list interface{}, view string) error {
	ret := m.Called()

	reflect.ValueOf(list).Elem().Set(reflect.ValueOf(ret.Get(0)))

	return nil
}

func (m *MockDB) Patch(table, id string, fields map[string]interface{}) error {
	return m.Called().Get(0).(error)
}
