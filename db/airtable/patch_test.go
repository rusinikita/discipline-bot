package airtable_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/stretchr/testify/assert"
)

// nolint:paralleltest // shoulg run after TestBase_One
func TestBase_Patch(t *testing.T) {
	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	var tasks []task.Task

	err = b.List(&tasks, db.Options{View: "TODO"})

	assert.NoError(t, err)

	err = b.Patch("Tasks", tasks[0].ID, map[string]interface{}{
		"Status": task.Done,
	})

	assert.NoError(t, err)

	err = b.Patch("Tasks", tasks[0].ID, map[string]interface{}{
		"Status": task.Todo,
	})

	assert.NoError(t, err)
}
