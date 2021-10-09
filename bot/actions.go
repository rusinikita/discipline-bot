package bot

import "gopkg.in/tucnak/telebot.v2"

// Actions to:
// send message
// send callback response
// edit message and it's buttons.
type action interface {
	do(b *telebot.Bot, r Request) error
}

type Request struct {
	m *telebot.Message
	c *telebot.Callback
}

type Message struct {
	Text string
}

func (m Message) do(b *telebot.Bot, r Request) error {
	_, err := b.Send(r.m.Sender, m.Text)

	return err
}

type Response struct {
	Text string
}

func (r Response) do(b *telebot.Bot, request Request) error {
	return b.Respond(request.c, &telebot.CallbackResponse{Text: r.Text})
}
