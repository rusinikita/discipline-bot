package db

type Base interface {
	One(table, id string, entity interface{}) error
	List(table string, list interface{}, view string) error
	Patch(table, id string, fields map[string]interface{}) error
}
