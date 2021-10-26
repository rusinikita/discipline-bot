package routine

import (
	"time"

	"github.com/rusinikita/discipline-bot/db"
)

type Routine struct {
	ID          db.ID
	Name        string
	Description string
	Trackers    []db.ID
}

type Try struct {
	Time    time.Time `json:",omitempty"`
	Routine db.ID
}

func (t Try) TableName() string {
	return "RoutineTries"
}
