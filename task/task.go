package task

import (
	"time"

	"github.com/rusinikita/discipline-bot/db"
)

type Task struct {
	ID            db.ID
	Name          string
	Business      *db.ID `json:",omitempty"`
	Note          string
	Condition     string `json:",omitempty"`
	Status        Status
	Created       time.Time `json:",omitempty"`
	StatusUpdated time.Time `json:",omitempty"`
}

type Status string

const (
	New      Status = ""
	Todo     Status = "Todo"
	Done     Status = "Done"
	Reflexed Status = "Reflexed"
)
