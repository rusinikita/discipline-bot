package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-multierror"
)

type Base struct {
	client *resty.Client
}

func New() (b Base, err error) {
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

	b.client = client

	return b, err
}
