package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/env"
	"github.com/rusinikita/discipline-bot/reminder"
	"github.com/rusinikita/discipline-bot/routine"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/rusinikita/discipline-bot/tracking"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not loaded")
	}

	db, err := airtable.New()
	if err != nil {
		log.Fatal(err)

		return
	}

	b, err := telebot.NewBot(settings())
	if err != nil {
		log.Fatal(err)

		return
	}

	handlers := append(
		task.Handlers(),
		tracking.Handlers()...,
	)
	handlers = append(handlers, routine.Handlers()...)

	startReminder := reminder.StartReminder{}
	handlers = append(handlers, startReminder, reminder.DebugReminder{})

	bot.RegisterHandlers(b, db, handlers)

	startReminder.Do(bot.DefaultBot(b, db))

	b.Start()
}

func settings() telebot.Settings {
	longPoller := &telebot.LongPoller{Timeout: 10 * time.Second}
	auth := func(u *telebot.Update) bool {
		switch {
		case u.Message != nil:
			return u.Message.Sender.ID == env.UserID()
		case u.Callback != nil:
			return u.Callback.Sender.ID == env.UserID()
		default:
			return false
		}
	}

	return telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: telebot.NewMiddlewarePoller(longPoller, auth),
	}
}
