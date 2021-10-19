package tracking

import (
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/tracking/record"
)

func Handlers() []bot.Handler {
	return []bot.Handler{
		record.Track{},
	}
}
