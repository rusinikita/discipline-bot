package routine

import (
	"fmt"
	"strings"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/tracking/tracker"
)

type start struct {
	Name string
}

func (s start) Description() string {
	return "Starts routine by name"
}

func (s start) Scan(data string) bot.Command {
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

	// get trackers
	var (
		routine  = routines[0]
		filter   = fmt.Sprintf(`FIND("%s", {Routines})`, routine.Name)
		trackers []tracker.Tracker
	)

	if b.Err(b.Base().List(&trackers, db.Options{Filter: filter})) {
		return
	}

	// todo: crate routine try

	// send message
	ts := make([]string, len(trackers))
	for i, t := range trackers {
		ts[i] = fmt.Sprintf("[ ] %s", t.Name)
	}

	text := fmt.Sprintf(
		"Start '%s' #routine\n\n%d/%d\n\n%s",
		routine.Name, 0, len(trackers),
		strings.Join(ts, "\n"),
	)

	// todo: add bool tracker buttons
	// todo: add input handling
	// todo: add force reply if has input trackers
	// todo: add other tracker buttons (to change reply input tracker)
	b.Action(bot.Message{
		Text:     text,
		DeleteRM: true,
	})
}
