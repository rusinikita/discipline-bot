package dbmock_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/db/dbmock"
	"github.com/stretchr/testify/assert"
)

func TestMockDB_List(t *testing.T) {
	t.Parallel()

	var (
		m      = dbmock.MockDB{}
		list   = []interface{}{1, "2", 3}
		result []interface{}
	)

	m.OnList().Return(list)

	err := m.List("", &result, "")

	assert.NoError(t, err)
	assert.Equal(t, list, result)
}
