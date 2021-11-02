package reminder

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/routine"
)

// nolint:gochecknoglobals // internal.
var stop = make(chan bool, 1)

type StartReminder struct{}

func (r StartReminder) Description() string {
	return "Restarts reminder to apply new time"
}

func (r StartReminder) Do(b bot.Bot) {
	stop <- true

	var routines []routine.Routine

	if b.Err(b.Base().List(&routines)) {
		return
	}

	gocron.Clear()

	for _, r := range routines {
		r := r
		if r.ReminderTime == "" {
			continue
		}

		err := gocron.Every(1).Day().At(r.ReminderTime).Do(func() {
			text := fmt.Sprintf("Time for '%s' #routine", r.Name)

			b.Action(bot.Message{
				Text: text,
				Buttons: []bot.Button{
					routine.NewStart(r.Name),
				},
			})
		})

		if b.Err(err) {
			fmt.Println(err)

			return
		}
	}

	_, time := gocron.NextRun()
	fmt.Printf(
		"Reminder initialized: jobs %d next run %s",
		len(gocron.Jobs()),
		time.Format("01/02/2006 15:04"),
	)

	stop = gocron.Start()
}
