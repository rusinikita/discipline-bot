package bot

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/rusinikita/discipline-bot/db"
	"gopkg.in/tucnak/telebot.v2"
)

type Handler interface {
	Do(Bot)
}

func Unique(h Handler) string {
	return strcase.ToSnake(reflect.ValueOf(h).Type().Name())
}

type Command interface {
	Description() string
	Do(b Bot)
}

func registerCommands(b *telebot.Bot, base db.Base, cc []Command) {
	menu := make([]telebot.Command, len(cc))

	for i, c := range cc {
		c := c
		endpoint := Unique(c)

		b.Handle("/"+endpoint, func(m *telebot.Message) {
			c.Do(fromMessage(b, m, base))
		})

		menu[i] = telebot.Command{
			Text:        endpoint,
			Description: c.Description(),
		}
	}

	logErr(b.SetCommands(menu))
}

type Button interface {
	Text() string
	Data() string
	Scan(data string) Button
	Do(b Bot)
}

func inlineButton(b Button) telebot.InlineButton {
	return telebot.InlineButton{
		Unique: Unique(b),
		Text:   b.Text(),
		Data:   b.Data(),
	}
}

func registerButtons(bot *telebot.Bot, base db.Base, bb []Button) {
	for _, b := range bb {
		b := b
		endpoint := "\f" + Unique(b)

		bot.Handle(endpoint, func(c *telebot.Callback) {
			b.Scan(c.Data).Do(fromCallback(bot, c, base))
		})
	}
}

func RegisterHandlers(b *telebot.Bot, db db.Base, hh []Handler) {
	var (
		commands []Command
		buttons  []Button
	)

	for _, h := range hh {
		switch handler := h.(type) {
		case Command:
			commands = append(commands, handler)
		case Button:
			buttons = append(buttons, handler)
		}
	}

	registerCommands(b, db, commands)
	registerButtons(b, db, buttons)
}
