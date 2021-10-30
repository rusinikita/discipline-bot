package env

import (
	"os"
	"strconv"
)

func Debug() bool {
	return os.Getenv("MODE") == "debug"
}

func UserID() int {
	s := os.Getenv("USER_ID")
	if s == "" {
		return 0
	}

	id, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return id
}
