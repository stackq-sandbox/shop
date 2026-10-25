package orders

import (
	"testing"

	"github.com/stackq-sandbox/shop/internal/cart"
	"github.com/stackq-sandbox/shop/internal/coupons"
)

func TestOrderEarnsCoupons(t *testing.T) {
	var c cart.Cart
	c.Add("sku-1", 2, 1999)
	c.Add("sku-2", 1, 499)
	o := Order{ID: "o-1", Cart: c}
	a, e, err := o.EarnCoupons(coupons.NewLedger(), "cust-1")
	if err != nil {
		t.Fatal(err)
	}
	if want := int64(o.Cart.Total() / 100); a.Points != want || e.Delta != want {
		t.Fatalf("points = %d, entry = %d, want %d", a.Points, e.Delta, want)
	}
}
