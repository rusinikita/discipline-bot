package bot

import (
	"fmt"
	"log"
	"reflect"

	"github.com/rusinikita/discipline-bot/db"
	"gopkg.in/tucnak/telebot.v2"
)

type Bot interface {
	// Do action or error
	Do(action interface{}) bool
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
			c: c,
		},
		db: db,
	}
}

func (b bot) Do(i interface{}) bool {
	switch action := i.(type) {
	case action:
		err := action.do(b.Bot, b.Request)

		logErr(err)

		return err != nil

	case error:
		return b.handleErr(action)
	default:
		return b.handleErr(fmt.Errorf(
			"unknown action: %s",
			reflect.TypeOf(i).Name(),
		))
	}
}

func (b bot) Base() db.Base {
	return b.db
}

func (b bot) handleErr(err error) bool {
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

func logErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
