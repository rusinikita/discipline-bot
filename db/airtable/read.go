package airtable

import (
	"errors"

	"github.com/rusinikita/discipline-bot/db"
)

func (b base) List(list interface{}, options ...db.Options) error {
	records := records{}

	request := b.client.R().SetResult(&records)

	if len(options) > 0 {
		o := options[0]

		if o.View != "" {
			request = request.SetQueryParam("view", o.View)
		}
	}

	r, err := request.Get(db.TableName(list))
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(string(r.Body()))
	}

	// mapping
	maps := make([]map[string]interface{}, len(records.Records))
	for i, r := range records.Records {
		maps[i] = r.Fields
		maps[i]["id"] = r.ID
	}

	return decode(maps, list)
}
