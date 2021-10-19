package routine

import "github.com/rusinikita/discipline-bot/db"

type Routine struct {
	ID          db.ID
	Name        string
	Description string
	Trackers    []db.ID
}
