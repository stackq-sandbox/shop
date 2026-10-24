package checkout

import "testing"

func order() Order {
	return Order{
		ID:      "o1",
		Lines:   []Line{{SKU: "a", Qty: 1, Unit: 10000}},
		Ship:    Address{Line1: "1 Main St", City: "San Francisco", Region: "CA", Postal: "94105", Country: "US"},
		Payment: "card",
	}
}

func TestPrice(t *testing.T) {
	q, err := Price(order())
	if err != nil {
		t.Fatal(err)
	}
	want := Quote{Subtotal: 10000, Discount: 0, Tax: 725, Shipping: 559, Total: 11284}
	if q != want {
		t.Fatalf("Price() = %+v, want %+v", q, want)
	}
	if got := q.Format(); got != "$112.84" {
		t.Fatalf("Format() = %q", got)
	}
}

func TestPriceCoupon(t *testing.T) {
	o := order()
	o.Coupon = "SAVE5"
	q, err := Price(o)
	if err != nil {
		t.Fatal(err)
	}
	if q.Discount != 500 {
		t.Fatalf("Discount = %d, want 500", q.Discount)
	}
	if q.Tax != 688 {
		t.Fatalf("Tax = %d, want 688 (7.25%% of 9500)", q.Tax)
	}
}

func TestPriceInternational(t *testing.T) {
	o := order()
	o.Ship = Address{Line1: "1 High St", City: "London", Postal: "SW1A 1AA", Country: "GB"}
	q, err := Price(o)
	if err != nil {
		t.Fatal(err)
	}
	if q.Tax != 0 {
		t.Fatalf("Tax = %d, want 0 outside the US", q.Tax)
	}
	if q.Shipping != 810 {
		t.Fatalf("Shipping = %d, want 810 (zone 4)", q.Shipping)
	}
}

func TestPriceRejectsBadOrders(t *testing.T) {
	empty := order()
	empty.Lines = nil
	if _, err := Price(empty); err == nil {
		t.Error("an order with no lines should fail")
	}
	bad := order()
	bad.Ship.Postal = "nope"
	if _, err := Price(bad); err == nil {
		t.Error("a bad postal code should fail")
	}
	region := order()
	region.Ship.Region = "ZZ"
	if _, err := Price(region); err == nil {
		t.Error("an unknown region should fail")
	}
}
