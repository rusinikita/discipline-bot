package db

import (
	"reflect"
	"strings"
)

type ID string

type Base interface {
	One(id string, entity interface{}) error
	List(list interface{}, options ...Options) error
	Create(entity interface{}) error
	Patch(table, id string, fields map[string]interface{}) error
	Delete(table, id string) error
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
	}

	return m
}
