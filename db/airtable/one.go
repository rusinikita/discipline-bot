package airtable

import "errors"

func (b base) One(table, id string, entity interface{}) error {
	record := record{}

	r, err := b.client.R().
		SetResult(&record).
		Get(table + "/" + id)
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	record.Fields["id"] = record.ID

	return decode(record.Fields, entity)
}
