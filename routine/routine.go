package routine

import (
	"fmt"
	"strings"
	"time"

	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db"
	"github.com/rusinikita/discipline-bot/tracking/record"
	"github.com/rusinikita/discipline-bot/tracking/tracker"
)

type Routine struct {
	ID           db.ID
	Name         string
	ReminderTime string
	Trackers     []db.ID
}

type Try struct {
	ID      db.ID
	Time    time.Time `json:",omitempty"`
	Routine db.ID
}

func (t Try) TableName() string {
	return "RoutineTries"
}

type try struct {
	Try
	Routine
	trackers []trackerState
}

type trackerState struct {
	tracker.Tracker
	tracked bool
}

func (t try) message() (m bot.Message) {
	ts := make([]string, len(t.Trackers))
	tracked := 0

	for i, track := range t.trackers {
		btn := NewTrackBoolBtn(track.Tracker, t.Try.ID, track.tracked)

		ts[i] = btn.Text()

		if track.tracked {
			tracked++
		}

		if track.Type == tracker.Bool {
			m.Buttons = append(m.Buttons, btn)
		}
	}

	m.Text = fmt.Sprintf(
		"Start '%s' #routine\n\n%d/%d\n\n%s",
		t.Routine.Name, tracked, len(t.Trackers),
		strings.Join(ts, "\n"),
	)

	// todo: add input handling
	// todo: add force reply if has input trackers
	// todo: add other tracker buttons (to change reply input tracker)
	return m
}

func getTry(id db.ID, b db.Base) (try try, err error) {
	if err = b.One(id, &try.Try); err != nil {
		return try, err
	}

	if err = b.One(try.Try.Routine, &try.Routine); err != nil {
		return try, err
	}

	var (
		tryPK          = try.Time.Format("01/02/2006 15:04")
		trackersFilter = fmt.Sprintf(`FIND("%s", {Routines})`, try.Name)
		recordsFilter  = fmt.Sprintf(`{RoutineTry} = '%s'`, tryPK)
		trackers       []tracker.Tracker
		records        []record.Record
	)

	if err = b.List(&trackers, db.Options{Filter: trackersFilter}); err != nil {
		return try, err
	}

	if err = b.List(&records, db.Options{Filter: recordsFilter}); err != nil {
		return try, err
	}

	for _, t := range trackers {
		state := trackerState{
			Tracker: t,
		}

		for _, r := range records {
			if r.Tracker == t.ID {
				state.tracked = true

				break
			}
		}

		try.trackers = append(try.trackers, state)
	}

	return try, err
}
