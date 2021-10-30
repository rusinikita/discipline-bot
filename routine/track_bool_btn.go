package routine

import (
	"fmt"
	"strings"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/tracking/record"
	"github.com/rusinikita/discipline-bot/tracking/tracker"
)

type trackBoolBtn struct {
	ID          db.ID
	TryID       db.ID
	Tracked     bool
	TrackerName string
}

func NewTrackBoolBtn(t tracker.Tracker, try db.ID, tracked bool) bot.Button {
	if t.Type != tracker.Bool {
		panic("only bool allowed")
	}

	return trackBoolBtn{
		ID:          t.ID,
		TryID:       try,
		Tracked:     tracked,
		TrackerName: t.Name,
	}
}

func (t trackBoolBtn) Text() string {
	emoji := "🔲"
	if t.Tracked {
		emoji = "☑️"
	}

	return fmt.Sprintf("%s %s", emoji, t.TrackerName)
}

func (t trackBoolBtn) Data() string {
	return fmt.Sprintf("%s,%s", t.ID, t.TryID)
}

func (t trackBoolBtn) Scan(data string) bot.Handler {
	ss := strings.Split(data, ",")

	t.ID, t.TryID = db.ID(ss[0]), db.ID(ss[1])

	return t
}

func (t trackBoolBtn) Do(b bot.Bot) {
	if t.Tracked {
		panic("not implemented")
	}

	if b.Err(b.Base().Create(record.NewBoolRecord(t.ID, t.TryID))) {
		return
	}

	try, err := getTry(t.TryID, b.Base())
	if b.Err(err) {
		return
	}

	m := try.message()

	b.Action(bot.Response{
		Text:        "done",
		EditMessage: &m,
	})
}
