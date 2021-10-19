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

func TestFields(t *testing.T) {
	t.Parallel()

	entity := task.Task{
		ID:       "1",
		Name:     "bla",
		Business: db.IDp("123"),
		Note:     "bla",
		Status:   task.Todo,
	}

	// no ID field
	// no Condition, Created, StatusUpdated field cause omitempty
	// no CreatedField cause computed in
	expected := map[string]interface{}{
		"Name":     entity.Name,
		"Business": []string{"123"},
		"Note":     entity.Note,
		"Status":   entity.Status,
	}

	assert.Equal(t, expected, db.Fields(entity))
}
