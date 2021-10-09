package airtable_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
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
	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	var tasks []task.Task

	err = b.List("Задачи", &tasks, "Tasks todo")

	log.Println(tasks)
	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)
}
