package discount

import (
	"testing"

	"github.com/stackq-sandbox/shop/internal/money"
)

func TestApply(t *testing.T) {
	cases := []struct {
		sub   money.Cents
		items int
		want  money.Cents
	}{
		{1000, 1, 1000},
		{1000, 2, 1000},
		{1000, 3, 900},
		{4497, 3, 4048},
	}
	for _, c := range cases {
		if got := Apply(c.sub, c.items); got != c.want {
			t.Errorf("Apply(%d, %d) = %d, want %d", c.sub, c.items, got, c.want)
		}
	}
}

func TestCoupon(t *testing.T) {
	cases := []struct {
		sub  money.Cents
		code string
		want money.Cents
	}{
		{10000, "SAVE5", 500},
		{300, "SAVE5", 300},
		{10000, "TENOFF", 1000},
		{10000, "HALF", 5000},
		{10000, "NOPE", 0},
		{10000, "", 0},
	}
	for _, c := range cases {
		if got := Coupon(c.sub, c.code); got != c.want {
			t.Errorf("Coupon(%d, %q) = %d, want %d", c.sub, c.code, got, c.want)
		}
	}
}
