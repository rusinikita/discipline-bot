package airtable

import (
	"errors"

	"github.com/rusinikita/discipline-bot/db"
)

func (b base) Create(entity interface{}) error {
	records := records{Records: []record{
		{Fields: db.Fields(entity)},
	}}

	r, err := b.client.R().
		SetBody(records).
		Post(db.TableName(entity))
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	return nil
}
