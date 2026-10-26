package solo

import "testing"

func TestName(t *testing.T) {
	if Name() != "solo" {
		t.Fatal("unexpected name")
	}
}
