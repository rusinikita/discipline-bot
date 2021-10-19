package airtable

import (
	"errors"

	"github.com/rusinikita/discipline-bot/db"
)

func (b base) Delete(table string, id db.ID) error {
	r, err := b.client.R().Delete(table + "/" + string(id))
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	return nil
}
