package task

import (
	"github.com/rusinikita/discipline-bot/bot"
)

// message reply -> bot send (sendable (any), replyMarkup ())
// action response -> update message + buttons, then respond to callback

const (
	tasksTable = "Tasks"
	todoView   = "TODO"
)

type CurrentTask struct{}

func (c CurrentTask) Description() string {
	return "Sends top task from TODO table view"
}

func (c CurrentTask) Do(b bot.Bot) {
	var tasks []Task

	if b.Err(b.Base().List(tasksTable, &tasks, todoView)) {
		return
	}

	t := tasks[0]

	b.Action(bot.Message{Text: currentTaskText(t.Name, t.Note)})
}
