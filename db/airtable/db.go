package airtable

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-multierror"
	"github.com/mitchellh/mapstructure"
	"github.com/rusinikita/discipline-bot/db"
)

type base struct {
	client *resty.Client
}

type records struct {
	Records []record `json:"records"`
}

type record struct {
	ID     db.ID                  `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields"`
}

func New() (b db.Base, err error) {
	id := os.Getenv("BASE_ID")
	if id == "" {
		err = multierror.Append(err, errors.New("BASE_ID env required"))
	}

	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		err = multierror.Append(err, errors.New("API_KEY env required"))
	}

	client := resty.New()

	client.SetDebug(os.Getenv("MODE") == "debug")
	// todo: setup airtable or nocodb
	client.SetHostURL("https://api.airtable.com/v0/" + id)
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(apikey)

	return base{client: client}, err
}

func decode(data interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  result,
		TagName: "json",
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			intToDuration,
			sliceToID,
		),
	})
	if err != nil {
		return err
	}

	return decoder.Decode(data)
}

func sliceToID(f, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.Slice {
		return data, nil
	}

	isPtr := t.Kind() == reflect.Ptr
	if isPtr {
		t = t.Elem()
	}

	if t != reflect.TypeOf(db.ID("")) {
		return data, nil
	}

	v := reflect.ValueOf(data)

	if v.Len() == 0 {
		if isPtr {
			return nil, nil
		}

		return data, errors.New("relation id required")
	}

	id := v.Index(0).Elem().String()
	if isPtr {
		return &id, nil
	}

	return db.ID(id), nil
}

func intToDuration(_, t reflect.Type, data interface{}) (interface{}, error) {
	if t != reflect.TypeOf(time.Duration(0)) {
		return data, nil
	}

	return time.ParseDuration(fmt.Sprint(data) + "s")
}
