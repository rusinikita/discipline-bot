package task

import "time"

type Task struct {
	ID            string
	Name          string
	Note          string
	Condition     string
	Status        Status
	Created       time.Time
	StatusUpdated time.Time
}

type Status string

const (
	New      Status = ""
	Todo     Status = "Todo"
	Done     Status = "Done"
	Reflexed Status = "Reflexed"
)
