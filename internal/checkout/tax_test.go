package checkout

import (
	"testing"

	"github.com/stackq-sandbox/shop/internal/money"
)

func TestTax(t *testing.T) {
	cases := []struct {
		region  string
		taxable money.Cents
		want    money.Cents
	}{
		{"AL", 10000, 400},
		{"AK", 10000, 0},
		{"AZ", 10000, 560},
		{"AR", 10000, 650},
		{"CA", 10000, 725},
		{"CO", 10000, 290},
		{"CT", 10000, 635},
		{"DE", 10000, 0},
		{"FL", 10000, 600},
		{"GA", 10000, 400},
		{"HI", 10000, 400},
		{"ID", 10000, 600},
		{"IL", 10000, 625},
		{"IN", 10000, 700},
		{"IA", 10000, 600},
		{"KS", 10000, 650},
		{"KY", 10000, 600},
		{"LA", 10000, 445},
		{"ME", 10000, 550},
		{"MD", 10000, 600},
		{"MA", 10000, 625},
		{"MI", 10000, 600},
		{"MN", 10000, 688},
		{"MS", 10000, 700},
		{"MO", 10000, 423},
		{"MT", 10000, 0},
		{"NE", 10000, 550},
		{"NV", 10000, 685},
		{"NH", 10000, 0},
		{"NJ", 10000, 663},
		{"NM", 10000, 513},
		{"NY", 10000, 400},
		{"NC", 10000, 475},
		{"ND", 10000, 500},
		{"OH", 10000, 575},
		{"OK", 10000, 450},
		{"OR", 10000, 0},
		{"PA", 10000, 600},
		{"RI", 10000, 700},
		{"SC", 10000, 600},
		{"SD", 10000, 450},
		{"TN", 10000, 700},
		{"TX", 10000, 625},
		{"UT", 10000, 485},
		{"VT", 10000, 600},
		{"VA", 10000, 430},
		{"WA", 10000, 650},
		{"WV", 10000, 600},
		{"WI", 10000, 500},
		{"WY", 10000, 400},
	}
	for _, c := range cases {
		got, err := Tax(c.region, c.taxable)
		if err != nil {
			t.Errorf("Tax(%q): %v", c.region, err)
			continue
		}
		if got != c.want {
			t.Errorf("Tax(%q, %d) = %d, want %d", c.region, c.taxable, got, c.want)
		}
	}
}

func TestTaxUnknownRegion(t *testing.T) {
	if _, err := Tax("ZZ", 100); err == nil {
		t.Fatal("expected an error for an unknown region")
	}
}
