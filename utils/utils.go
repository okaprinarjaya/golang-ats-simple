package utils

import (
	"time"
)

func DateValid(date time.Time) bool {
	if date.IsZero() {
		return false
	} else {
		return true
	}
}

func StringValid(str string) bool {
	if str == "" {
		return false
	} else {
		return true
	}
}
