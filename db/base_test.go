package db_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Tasks", db.TableName(&task.Task{}))
	assert.Equal(t, "Tasks", db.TableName([]task.Task{}))
	assert.Equal(t, "Tasks", db.TableName(&[]task.Task{}))
}
