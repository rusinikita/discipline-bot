package routine

import "github.com/rusinikita/discipline-bot/bot"

func Handlers() []bot.Handler {
	return []bot.Handler{start{}, trackBoolBtn{}}
}
