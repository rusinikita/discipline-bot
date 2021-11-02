package bot

import (
	"log"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/rusinikita/discipline-bot/db"
	"gopkg.in/tucnak/telebot.v2"
)

type Handler interface {
	Do(Bot)
}

type Scanner interface {
	Scan(data string) Handler
}

func Unique(h Handler) string {
	return strcase.ToSnake(reflect.ValueOf(h).Type().Name())
}

type Command interface {
	Description() string
	Handler
}

func registerCommands(b *telebot.Bot, base db.Base, cc []Command) {
	menu := make([]telebot.Command, len(cc))

	for i, c := range cc {
		c := c
		endpoint := Unique(c)

		b.Handle("/"+endpoint, func(m *telebot.Message) {
			b := fromMessage(b, m, base)

			if s, ok := c.(Scanner); ok {
				//nolint:forcetypeassert // 100% command
				c = s.Scan(strings.TrimSpace(m.Payload)).(Command)
			}

			c.Do(b)
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
	Handler
	Scanner
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
		if c, ok := h.(Command); ok {
			commands = append(commands, c)
		}

		if b, ok := h.(Button); ok {
			buttons = append(buttons, b)
		}

		switch h.(type) {
		case Command, Button:
			continue
		default:
			log.Fatalf("unknown handler type: %T", h)
		}
	}

	registerCommands(b, db, commands)
	registerButtons(b, db, buttons)
}
