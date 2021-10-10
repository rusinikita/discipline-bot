package bot

import (
	"log"

	"github.com/rusinikita/discipline-bot/db"
	"gopkg.in/tucnak/telebot.v2"
)

type Bot interface {
	// Do Action or error
	Action(action Action)
	Err(err error) bool
	Base() db.Base
}

type bot struct {
	*telebot.Bot
	Request
	db db.Base
}

func fromMessage(b *telebot.Bot, m *telebot.Message, db db.Base) Bot {
	return bot{
		Bot: b,
		Request: Request{
			m: m,
		},
		db: db,
	}
}

func fromCallback(b *telebot.Bot, c *telebot.Callback, db db.Base) Bot {
	return bot{
		Bot: b,
		Request: Request{
			m: c.Message,
			c: c,
		},
		db: db,
	}
}

func (b bot) Action(action Action) {
	err := action.Do(b.Bot, b.Request)

	logErr(err)
}

func (b bot) Err(err error) bool {
	if err == nil {
		return false
	}

	logErr(err)

	if c := b.Request.c; c != nil {
		logErr(b.Bot.Respond(c, &telebot.CallbackResponse{
			Text:      err.Error(),
			ShowAlert: true,
		}))
	}

	if m := b.Request.m; m != nil {
		_, err = b.Bot.Send(m.Sender, err.Error())

		logErr(err)
	}

	return true
}

func (b bot) Base() db.Base {
	return b.db
}

func logErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
