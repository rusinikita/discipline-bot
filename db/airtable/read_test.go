package airtable_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/task"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestBase_List(t *testing.T) {
	t.Parallel()

	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	var tasks []task.Task

	t.Run("list view", func(t *testing.T) {
		t.Parallel()

		err = b.List(&tasks, db.Options{View: "TODO"})

		assert.NoError(t, err)
		assert.NotEmpty(t, tasks)

		for i, _t := range tasks {
			assert.NotEmpty(t, _t.Name, "%d %s was empty", i, _t.ID)
			assert.Equal(t, _t.Status, task.Todo)
		}
	})

	t.Run("list filter", func(t *testing.T) {
		t.Parallel()

		name := "Test task"

		err = b.List(&tasks, db.Options{Filter: task.Task{Name: name}})

		assert.NoError(t, err)
		assert.NotEmpty(t, tasks)
		assert.Equal(t, name, tasks[0].Name)
	})
}
