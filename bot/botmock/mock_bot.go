package botmock

import (
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/dbmock"
	"github.com/stretchr/testify/mock"
)

type BotMock struct {
	mock.Mock
	db db.Base
}

func (b *BotMock) Do(action interface{}) bool {
	return b.Called(action).Get(0).(bool)
}

func (b *BotMock) Base() db.Base {
	if b.db == nil {
		b.db = &dbmock.MockDB{}
	}

	return b.db
}
