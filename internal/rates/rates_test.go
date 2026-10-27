package rates

import "testing"

func TestFor(t *testing.T) {
	q := For(1000)
	if len(q) != 2 || q[0].Cents != 400 || q[1].Cents != 800 {
		t.Fatalf("unexpected quotes: %+v", q)
	}
}
