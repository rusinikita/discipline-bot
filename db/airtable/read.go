package airtable

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

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

		if o.Filter != nil {
			request = request.SetQueryParam("filterByFormula", filterString(o.Filter))
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

func filterString(filter interface{}) string {
	if s, ok := filter.(string); ok {
		return s
	}

	var fieldFilters []string //nolint:prealloc // can't preallocate

	for key, value := range db.Fields(filter) {
		if reflect.ValueOf(value).IsZero() {
			continue
		}

		// potential bug with '' wrapping
		fieldFilter := fmt.Sprintf("{%s} = '%v'", key, value)

		fieldFilters = append(fieldFilters, fieldFilter)
	}

	return fmt.Sprintf("AND(%s)", strings.Join(fieldFilters, ", "))
}
