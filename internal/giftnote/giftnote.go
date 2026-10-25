// Package giftnote validates the note a customer adds to a gift order.
package giftnote

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// MaxLen is the longest note, in characters.
const MaxLen = 200

// ErrTooLong means the note has more than MaxLen characters.
var ErrTooLong = errors.New("giftnote: note is too long")

// Clean trims the note and checks its length.
func Clean(note string) (string, error) {
	note = strings.TrimSpace(note)
	if utf8.RuneCountInString(note) > MaxLen {
		return "", ErrTooLong
	}
	return note, nil
}
