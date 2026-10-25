package giftnote

import (
	"errors"
	"strings"
	"testing"
)

func TestClean(t *testing.T) {
	if got, err := Clean("  Happy birthday!  "); err != nil || got != "Happy birthday!" {
		t.Fatalf("Clean = %q, %v", got, err)
	}
	if _, err := Clean(strings.Repeat("é", MaxLen+1)); !errors.Is(err, ErrTooLong) {
		t.Fatalf("err = %v, want ErrTooLong", err)
	}
}
