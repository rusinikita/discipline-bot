package airtable_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/stretchr/testify/assert"
)

func TestBase_One(t *testing.T) {
	t.Parallel()

	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	var tasks []task.Task

	err = b.List(&tasks, db.Options{View: "TODO"})

	assert.NoError(t, err)

	result := task.Task{}

	err = b.One(tasks[0].ID, &result)

	assert.NoError(t, err)
	assert.Equal(t, tasks[0], result)
}
