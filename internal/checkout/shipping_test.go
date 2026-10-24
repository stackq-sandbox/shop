package checkout

import (
	"testing"

	"github.com/stackq-sandbox/shop/internal/money"
)

func TestShipping(t *testing.T) {
	cases := []struct {
		carrier string
		zone    int
		grams   int
		want    money.Cents
	}{
		{"ground", 1, 500, 559},
		{"ground", 2, 500, 642},
		{"express", 1, 1000, 1539},
		{"overnight", 1, 0, 2499},
		{"economy", 8, 2000, 1084},
		{"pickup", 5, 9000, 0},
	}
	for _, c := range cases {
		got, err := Shipping(c.carrier, c.zone, c.grams)
		if err != nil {
			t.Errorf("Shipping(%q, %d, %d): %v", c.carrier, c.zone, c.grams, err)
			continue
		}
		if got != c.want {
			t.Errorf("Shipping(%q, %d, %d) = %d, want %d", c.carrier, c.zone, c.grams, got, c.want)
		}
	}
}

func TestShippingErrors(t *testing.T) {
	if _, err := Shipping("pigeon", 1, 1); err == nil {
		t.Error("unknown carrier should fail")
	}
	if _, err := Shipping("ground", 0, 1); err == nil {
		t.Error("zone 0 should fail")
	}
	if _, err := Shipping("ground", 9, 1); err == nil {
		t.Error("zone 9 should fail")
	}
	if _, err := Shipping("ground", 1, -1); err == nil {
		t.Error("negative weight should fail")
	}
}

func TestZone(t *testing.T) {
	for country, want := range map[string]int{"US": 1, "CA": 2, "DE": 4, "SE": 5, "JP": 7, "BR": 8} {
		if got := Zone(country); got != want {
			t.Errorf("Zone(%q) = %d, want %d", country, got, want)
		}
	}
}
