// Package discount holds the volume and coupon discounts for the new checkout.
package discount

import "github.com/stackq-sandbox/shop/internal/money"

// Apply takes 10% off orders of three or more items.
func Apply(sub money.Cents, items int) money.Cents {
	if items >= 3 {
		return sub - sub/10
	}
	return sub
}

// Coupon returns the amount a coupon code takes off sub. Unknown codes are
// worth nothing.
func Coupon(sub money.Cents, code string) money.Cents {
	var off money.Cents
	switch code {
	case "SAVE5":
		off = 500
	case "TENOFF":
		off = sub / 10
	case "HALF":
		off = sub / 2
	default:
		return 0
	}
	if off > sub {
		return sub
	}
	return off
}
