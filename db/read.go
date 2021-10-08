package db

import "errors"

type records struct {
	Records interface{}
}

func (b Base) List(table string, list interface{}, view string) error {
	records := records{
		Records: list,
	}

	r, err := b.client.R().
		SetPathParam("view", view).
		SetResult(&records).
		Get(table)
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	return err
}
