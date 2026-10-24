package orders

import "testing"

func TestSummary(t *testing.T) {
	o := Order{ID: "o1"}
	o.Cart.Add("a", 2, 1250)
	if got, want := o.Summary(), "o1: $25.00"; got != want {
		t.Fatalf("Summary() = %q, want %q", got, want)
	}
}
