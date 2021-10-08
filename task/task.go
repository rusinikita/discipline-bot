package task

import "time"

type Task struct {
	ID            string
	Name          string
	Note          string `json:"Notes"`
	Condition     string
	Status        Status
	Created       time.Time `json:"created_at"`
	StatusUpdated time.Time `json:"status_updated_at"`
}

type Status string

const (
	New      Status = ""
	Todo     Status = "Todo"
	Done     Status = "Done"
	Reflexed Status = "Reflexed"
)
