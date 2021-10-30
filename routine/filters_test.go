package routine_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/routine"
	"github.com/rusinikita/discipline-bot/tracking/record"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestFilters(t *testing.T) {
	t.Parallel()

	b, err := airtable.New()
	if err != nil {
		t.Fatal(err)
	}

	try := routine.Try{}

	assert.NoError(t, b.One("rec3rWY6c5V2Vw6SX", &try))
	assert.NotEmpty(t, try.ID)

	var (
		filter  = fmt.Sprintf(`{RoutineTry} = '%s'`, try.Time.Format("01/02/2006 15:04"))
		records []record.Record
	)

	assert.NoError(t, b.List(&records, db.Options{Filter: filter}))
	assert.NotEmpty(t, records)
}
