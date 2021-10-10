package airtable

import "errors"

func (b base) Patch(table, id string, fields map[string]interface{}) error {
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
