package db

type Base interface {
	List(table string, list interface{}, view string) error
}
