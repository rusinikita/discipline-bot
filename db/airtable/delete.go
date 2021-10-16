package airtable

import (
	"errors"
)

func (b base) Delete(table, id string) error {
	r, err := b.client.R().Delete(table + "/" + id)
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	return nil
}
