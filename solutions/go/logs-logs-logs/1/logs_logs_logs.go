package logs

import (
	"fmt"
	"unicode/utf8"
)

func Application(log string) string {
	for _, v := range log {
		if fmt.Sprintf("%U", v) == "U+2757" {
			return "recommendation"
		} else if fmt.Sprintf("%U", v) == "U+1F50D" {
			return "search"
		} else if fmt.Sprintf("%U", v) == "U+2600" {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, v := range log {
		if v == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(v)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
