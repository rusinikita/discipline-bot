package task

import (
	"github.com/rusinikita/discipline-bot/bot"
)

type taskDone struct {
	taskID string
}

func (d taskDone) Text() string {
	return done
}

func (d taskDone) Data() string {
	return d.taskID
}

func (d taskDone) Scan(data string) bot.Button {
	d.taskID = data

	return d
}

func (d taskDone) Do(b bot.Bot) {
	t := Task{}

	if b.Err(b.Base().One(tasksTable, d.taskID, &t)) {
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
