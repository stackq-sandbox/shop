package cart

import "testing"

func TestTotal(t *testing.T) {
	var c Cart
	c.Add("a", 1, 1000)
	c.Add("b", 1, 500)
	if got := c.Total(); got != 1500 {
		t.Fatalf("Total() = %d, want 1500", got)
	}
	c.Add("a", 1, 1000)
	if got := c.Total(); got != 2250 {
		t.Fatalf("Total() with volume discount = %d, want 2250", got)
	}
}
