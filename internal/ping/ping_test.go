package ping

import "testing"

func TestCheck(t *testing.T) {
	if Check() != Reply {
		t.Fatal("unexpected reply")
	}
}
