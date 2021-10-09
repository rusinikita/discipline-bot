package botmock_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/bot/botmock"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/bot/botmock"
	"github.com/stretchr/testify/mock"
)

func TestBotMock_Do(t *testing.T) {
	m := botmock.BotMock{}

	message := bot.Message{Text: "123"}
	response := bot.Response{Text: "done"}

	m.On("Do", mock.Anything).Return(true)

	m.Do(message)
	m.Do(response)

	m.AssertCalled(t, "Do", message)
	m.AssertCalled(t, "Do", response)
}
