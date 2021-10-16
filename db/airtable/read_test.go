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

	err = b.List(&tasks, db.Options{View: "TODO"})

	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)

	for i, _t := range tasks {
		assert.NotEmpty(t, _t.Name, "%d %s was empty", i, _t.ID)
		assert.Equal(t, _t.Status, task.Todo)
	}
}
