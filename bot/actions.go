package bot

import (
	"github.com/rusinikita/discipline-bot/env"
	"gopkg.in/tucnak/telebot.v2"
)

// Actions to:
// send message
// send callback response
// edit message and it's buttons.
type Action interface {
	Do(b *telebot.Bot, r Request) error
}

type Request struct {
	m *telebot.Message
	c *telebot.Callback
}

func (r Request) user() *telebot.User {
	if r.c != nil {
		return r.c.Sender
	}

	if r.m != nil {
		return r.m.Sender
	}

	if id := env.UserID(); id != 0 {
		return &telebot.User{ID: id}
	}

	panic("no user source")
}

type Message struct {
	Text     string
	DeleteRM bool // Deletes request message from chat history
	Buttons  []Button
}

func (m Message) options() *telebot.SendOptions {
	keyboard := make([][]telebot.InlineButton, len(m.Buttons))

	for i := range m.Buttons {
		keyboard[i] = []telebot.InlineButton{inlineButton(m.Buttons[i])}
	}

	return &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
		ReplyMarkup: &telebot.ReplyMarkup{
			InlineKeyboard: keyboard,
		},
	}
}

func (m Message) Do(b *telebot.Bot, r Request) error {
	_, err := b.Send(r.user(), m.Text, m.options())
	if err != nil {
		return err
	}

	if m.DeleteRM {
		return b.Delete(r.m)
	}

	return nil
}

type Response struct {
	Text        string
	EditMessage *Message
}

func (r Response) Do(b *telebot.Bot, request Request) error {
	if m := r.EditMessage; m != nil {
		_, err := b.Edit(request.m, m.Text, m.options())
		if err != nil {
			return err
		}
	}

	return b.Respond(request.c, &telebot.CallbackResponse{Text: r.Text})
}
