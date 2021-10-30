package record

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/tracking/tracker"
)

type Track struct {
	TrackerName string
	Text        string
}

func (t Track) Scan(data string) bot.Handler {
	ss := make([]string, 2)

	for i, s := range strings.SplitN(data, " ", 2) {
		ss[i] = s
	}

	t.TrackerName = strings.TrimSpace(ss[0])
	t.Text = strings.TrimSpace(ss[1])

	return t
}

func (t Track) Description() string {
	return "creates new tracker record"
}

func (t Track) Do(b bot.Bot) {
	if b.Err(empty("tracker name", t.TrackerName)) {
		return
	}

	filter := tracker.Tracker{Name: t.TrackerName}

	var trackers []tracker.Tracker

	if b.Err(b.Base().List(&trackers, db.Options{Filter: filter})) {
		return
	} else if len(trackers) == 0 {
		b.Err(errors.New("no trackers found"))

		return
	}

	_tracker := trackers[0]

	newRecord, err := NewRecord(t.Text, _tracker)
	if b.Err(err) {
		return
	}

	if b.Err(b.Base().Create(newRecord)) {
		return
	}

	b.Action(bot.Message{
		Text:     recordCreated(_tracker, t.Text),
		DeleteRM: true,
	})
}

func empty(name, value string) error {
	if value != "" {
		return nil
	}

	return fmt.Errorf("%s required", name)
}
