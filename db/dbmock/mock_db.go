package dbmock

import (
	"reflect"

	"github.com/rusinikita/discipline-bot/db"
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

func (m *MockDB) OnCreate() *mock.Call {
	return m.On("Create")
}

func (m *MockDB) OnPatch() *mock.Call {
	return m.On("Patch")
}

func (m *MockDB) OnDelete() *mock.Call {
	return m.On("Delete")
}

func (m *MockDB) One(id db.ID, entity interface{}) error {
	return m.Called().Get(0).(error)
}

func (m *MockDB) List(list interface{}, options ...db.Options) error {
	ret := m.Called()

	reflect.ValueOf(list).Elem().Set(reflect.ValueOf(ret.Get(0)))

	return nil
}

func (m *MockDB) Create(entity interface{}) error {
	return m.Called().Get(0).(error)
}

func (m *MockDB) Patch(table string, id db.ID, fields map[string]interface{}) error {
	return m.Called().Get(0).(error)
}

func (m *MockDB) Delete(table string, id db.ID) error {
	return m.Called().Get(0).(error)
}
