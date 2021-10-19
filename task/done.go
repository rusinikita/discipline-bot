package task

import (
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
)

type taskDone struct {
	taskID db.ID
}

func (d taskDone) Text() string {
	return done
}

func (d taskDone) Data() string {
	return string(d.taskID)
}

func (d taskDone) Scan(data string) bot.Button {
	d.taskID = db.ID(data)

	return d
}

func (d taskDone) Do(b bot.Bot) {
	t := Task{}

	if b.Err(b.Base().One(d.taskID, &t)) {
		return
	}

	fields := map[string]interface{}{
		"Status": Done,
	}

	if b.Err(b.Base().Patch(tasksTable, d.taskID, fields)) {
		return
	}

	b.Action(bot.Response{
		Text: bot.Done,
		EditMessage: &bot.Message{
			Text: doneTaskText(t.Name),
		},
	})

	currentTask{}.Do(b)
}
