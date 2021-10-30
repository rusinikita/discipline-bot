package reminder

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/rusinikita/discipline-bot/bot"
)

type DebugReminder struct{}

func (d DebugReminder) Description() string {
	return "returns debug info"
}

func (d DebugReminder) Do(b bot.Bot) {
	jobTime := ""
	j, time := gocron.NextRun()

	if j != nil {
		jobTime = j.NextScheduledTime().Format("01/02/2006 15:04")
	}

	t := fmt.Sprintf(
		"jobs %d, next run %s, job time %s",
		len(gocron.Jobs()),
		time.Format("01/02/2006 15:04"),
		jobTime,
	)

	b.Action(bot.Message{
		Text:     t,
		DeleteRM: true,
	})
}
