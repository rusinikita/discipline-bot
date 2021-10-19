package record

import (
	"fmt"

	"github.com/rusinikita/discipline-bot/tracking/tracker"
)

const (
	recordCreatedText = "#tracked %s for %s"
)

func recordCreated(tracker tracker.Tracker, input string) string {
	return fmt.Sprintf(recordCreatedText, input, tracker.Name)
}
