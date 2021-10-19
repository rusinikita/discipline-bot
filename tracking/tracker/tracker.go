package tracker

import "github.com/rusinikita/discipline-bot/db"

type Tracker struct {
	ID          db.ID
	Name        string
	Description string
	Type
}

type Type string

const (
	Bool     Type = "Bool"
	Number   Type = "Number"
	Rating   Type = "Rating"
	Text     Type = "Text"
	Duration Type = "Duration"
)
