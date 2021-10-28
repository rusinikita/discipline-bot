package db

import (
	"reflect"
	"strings"
)

type ID string

func IDp(s string) *ID {
	id := ID(s)

	return &id
}

type Base interface {
	One(id ID, entity interface{}) error
	List(list interface{}, options ...Options) error
	Create(entity interface{}) error
	Patch(table string, id ID, fields map[string]interface{}) error
	Delete(table string, id ID) error
}

type TableNamer interface {
	TableName() string
}

type Options struct {
	View   string      // view for records filter
	Filter interface{} // entity obj with not zero values filter fields
}

func TableName(entity interface{}) string {
	t := reflect.ValueOf(entity).Type()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	if namer, ok := reflect.New(t).Interface().(TableNamer); ok {
		return namer.TableName()
	}

	return t.Name() + "s"
}

func Fields(entity interface{}) map[string]interface{} {
	v := reflect.ValueOf(entity)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("entity must be struct type or ptr")
	}

	m := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		fType := v.Type().Field(i)
		fName := fType.Name
		fValue := v.Field(i)

		if fName == "ID" {
			continue
		}

		omitempty := strings.Contains(fType.Tag.Get("json"), "omitempty")
		if omitempty && fValue.IsZero() {
			continue
		}

		m[fName] = fValue.Interface()

		if fType.Type.Kind() == reflect.Ptr {
			fType.Type = fType.Type.Elem()
			fValue = fValue.Elem()
		}

		if fType.Type == reflect.TypeOf(ID("")) {
			m[fName] = []string{fValue.String()}
		}
	}

	return m
}

func SetID(id ID, entity interface{}) {
	v := reflect.ValueOf(entity)
	if v.Kind() != reflect.Ptr {
		return
	}

	v.Elem().FieldByName("ID").Set(reflect.ValueOf(id))
}
