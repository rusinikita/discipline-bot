package airtable_test

import (
	"testing"

	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/stretchr/testify/assert"
)

func TestBase_Create(t *testing.T) {
	t.Parallel()

	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	var tasks []task.Task

	update := func() {
		tasks = nil

		err = b.List(&tasks, db.Options{View: "TODO"})
		if err != nil {
			t.Fatalf("list err %s", err)
		}
	}

	update()

	prevLen := len(tasks)

	err = b.Create(task.Task{
		Name:   "Bla",
		Note:   "Bla",
		Status: task.Todo,
	})

	assert.NoError(t, err)

	update()
	assert.Len(t, tasks, prevLen+1)
	lastTask := tasks[len(tasks)-1]
	assert.Equal(t, "Bla", lastTask.Name)

	// don't run delete
	if len(tasks) == prevLen {
		t.Skip("nothing created to delete")
	}

	err = b.Delete(db.TableName(lastTask), lastTask.ID)
	assert.NoError(t, err)

	update()
	assert.Len(t, tasks, prevLen)
}
