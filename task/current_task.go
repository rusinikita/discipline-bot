package task

import (
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
)

// message reply -> bot send (sendable (any), replyMarkup ())
// action response -> update message + buttons, then respond to callback

const (
	tasksTable = "Tasks"
	todoView   = "TODO"
)

type currentTask struct{}

func (c currentTask) Description() string {
	return "Sends top task from TODO table view"
}

func (c currentTask) Do(b bot.Bot) {
	var tasks []Task

	if b.Err(b.Base().List(&tasks, db.Options{View: todoView})) {
		return
	}

	t := tasks[0]

	b.Action(bot.Message{
		Text:    currentTaskText(t.Name, t.Note),
		Buttons: []bot.Button{taskDone{taskID: t.ID}},
	})
}
