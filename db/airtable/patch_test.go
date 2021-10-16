package airtable_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/stretchr/testify/assert"
)

// nolint:paralleltest // should run after TestBase_One
func TestBase_Patch(t *testing.T) {
	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	var tasks []task.Task

	err = b.List(&tasks, db.Options{View: "TODO"})

	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)
	assert.NotEqual(t, tasks[0].Status, task.Done)

	err = b.Patch("Tasks", tasks[0].ID, map[string]interface{}{
		"Status": task.Done,
	})

	assert.NoError(t, err)

	expected := task.Task{}

	err = b.One(tasks[0].ID, &expected)

	assert.NoError(t, err)
	assert.Equal(t, expected.Status, task.Done)

	// return state
	err = b.Patch("Tasks", tasks[0].ID, map[string]interface{}{
		"Status": task.Todo,
	})

	assert.NoError(t, err)
}
