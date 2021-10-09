package botmock

import (
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/dbmock"
	"github.com/stretchr/testify/mock"
)

type BotMock struct {
	mock.Mock
	db db.Base
}

func (b *BotMock) Action(action bot.Action) {
	b.Called(action)
}

func (b *BotMock) Err(err error) bool {
	return b.Called(err).Get(0).(bool)
}

func (b *BotMock) Base() db.Base {
	if b.db == nil {
		b.db = &dbmock.MockDB{}
	}

	return b.db
}
