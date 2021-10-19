package record

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/tracking/tracker"
)

type Record struct {
	ID       db.ID
	Time     time.Time     `json:",omitempty"`
	Tracker  db.ID         `json:",omitempty"`
	Rating   uint          `json:",omitempty"`
	Number   uint          `json:",omitempty"`
	Duration time.Duration `json:",omitempty"`
	Text     string        `json:",omitempty"`
}

func NewRecord(text string, parent tracker.Tracker) (r Record, err error) {
	r.Tracker = parent.ID

	integer, err := strconv.Atoi(text)

	switch parent.Type {
	case tracker.Text:
		r.Text = text
	case tracker.Bool:
	case tracker.Rating:
		if err != nil {
			return r, err
		}

		if integer > 10 || integer < 1 {
			err = errors.New("rating must in 1..10")

			return r, err
		}

		r.Rating = uint(integer)
	case tracker.Number:
		if err != nil {
			return r, err
		}

		r.Number = uint(integer)
	case tracker.Duration:
		return r, fmt.Errorf("%s type not implemented", parent.Type)
	}

	return r, nil
}
