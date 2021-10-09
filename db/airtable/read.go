package airtable

import (
	"errors"
	"time"

	"github.com/mitchellh/mapstructure"
)

type records struct {
	Records []record
}

type record struct {
	ID     string
	Fields map[string]interface{}
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

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     list,
		TagName:    "json",
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
	})
	if err != nil {
		return err
	}

	err = decoder.Decode(maps)
	if err != nil {
		return err
	}

	return err
}
