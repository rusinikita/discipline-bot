package task

import "time"

type Task struct {
	ID            string
	Name          string
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
