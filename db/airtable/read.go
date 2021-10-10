package airtable

import (
	"errors"
)

type records struct {
	Records []record `json:"records"`
}

type record struct {
	ID     string                 `json:"id"`
	Fields map[string]interface{} `json:"fields"`
}

func (b base) List(table string, list interface{}, view string) error {
	records := records{}

	r, err := b.client.R().
		SetQueryParam("view", view).
		SetResult(&records).
		Get(table)
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
