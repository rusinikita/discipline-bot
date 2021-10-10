package task

import "fmt"

const (
	done              = "Готово"
	currentTaskFormat = "Current task:  <b>%s</b>\n\n%s"
	doneTaskFormat    = "Task done: <b>%s</b>"
)

func currentTaskText(name, notes string) string {
	return fmt.Sprintf(currentTaskFormat, name, notes)
}

func doneTaskText(name string) string {
	return fmt.Sprintf(doneTaskFormat, name)
}
