package airtable

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-multierror"
	"github.com/mitchellh/mapstructure"
	"github.com/rusinikita/discipline-bot/db"
)

type base struct {
	client *resty.Client
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

	// todo: setup airtable or nocodb
	client.SetHostURL("https://api.airtable.com/v0/" + id)
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(apikey)
	client.OnResponseLog(func(log *resty.ResponseLog) error {
		_, err := fmt.Println(*log)

		return err
	})

	return base{client: client}, err
}

func decode(data interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     result,
		TagName:    "json",
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
	})
	if err != nil {
		return err
	}

	return decoder.Decode(data)
}
