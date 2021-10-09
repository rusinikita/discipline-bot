package dbtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockDB_List(t *testing.T) {
	var (
		m      = MockDB{}
		list   = []interface{}{1, "2", 3}
		result []interface{}
	)

	m.OnList().Return(list)

	err := m.List("", &result, "")

	assert.NoError(t, err)
	assert.Equal(t, list, result)
}
