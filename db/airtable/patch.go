package airtable

import (
	"errors"

	"github.com/rusinikita/discipline-bot/db"
)

func (b base) Patch(table string, id db.ID, fields map[string]interface{}) error {
	r, err := b.client.R().
		SetBody(records{Records: []record{
			{
				ID:     id,
				Fields: fields,
			},
		}}).
		Patch(table)
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	return nil
}
