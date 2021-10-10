package task

import "github.com/rusinikita/discipline-bot/bot"

func Handlers() []bot.Handler {
	return []bot.Handler{
		currentTask{},
		taskDone{},
	}
}
