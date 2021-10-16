package db

import "reflect"

type ID string

type Base interface {
	One(id string, entity interface{}) error
	List(list interface{}, options ...Options) error
	Create(entity interface{}) error
	Patch(table, id string, fields map[string]interface{}) error
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
