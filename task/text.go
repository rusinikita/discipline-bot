package task

import "fmt"

const (
	currentTaskFormat = "Current task: <b>%s</b>\n\n%s"
)

func currentTaskText(name, notes string) string {
	return fmt.Sprintf(currentTaskFormat, name, notes)
}
