package routine

import (
	"fmt"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
)

type start struct {
	Name string
}

func (s start) Text() string {
	return "Start"
}

func (s start) Data() string {
	return s.Name
}

func (s start) Description() string {
	return "Starts routine by name"
}

func (s start) Scan(data string) bot.Handler {
	s.Name = data

	return s
}

func (s start) Do(b bot.Bot) {
	// get routine
	var routines []Routine

	if b.Err(b.Base().List(&routines, db.Options{Filter: Routine{Name: s.Name}})) {
		return
	}

	if len(routines) == 0 {
		b.Action(bot.Message{
			Text: fmt.Sprintf("Routine '%s' not found", s.Name),
		})

		return
	}

	// create try
	newTry := Try{Routine: routines[0].ID}

	if b.Err(b.Base().Create(&newTry)) {
		return
	}

	// return state message
	t, err := getTry(newTry.ID, b.Base())
	if b.Err(err) {
		return
	}

	m := t.message()

	m.DeleteRM = true

	b.Action(m)
}
