package botmock_test

import (
	"errors"
	"testing"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/bot/botmock"
	"github.com/stretchr/testify/mock"
)

func TestBotMock_Do(t *testing.T) {
	t.Parallel()

	m := botmock.BotMock{}

	message := bot.Message{Text: "123"}
	response := bot.Response{Text: "done"}
	err := errors.New("test")

	m.On("Action", mock.Anything)
	m.On("Err", mock.Anything).Return(true)

	m.Action(message)
	m.Action(response)
	m.Err(err)

	m.AssertCalled(t, "Action", message)
	m.AssertCalled(t, "Action", response)
	m.AssertCalled(t, "Err", err)
}
