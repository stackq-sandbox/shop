package checkout

import (
	"fmt"
	"strings"

	"github.com/stackq-sandbox/shop/internal/checkout/discount"
	"github.com/stackq-sandbox/shop/internal/money"
)

// gramsPerItem is the flat parcel weight assumed per item.
const gramsPerItem = 500

// Price validates o and returns what the customer would pay.
func Price(o Order) (Quote, error) {
	if len(o.Lines) == 0 {
		return Quote{}, fmt.Errorf("checkout: order %q has no lines", o.ID)
	}
	if err := ValidateAddress(o.Ship); err != nil {
		return Quote{}, err
	}
	sub := o.Subtotal()
	after := discount.Apply(sub, o.Items())
	after -= discount.Coupon(after, o.Coupon)
	q := Quote{Subtotal: sub, Discount: sub - after}

	if strings.EqualFold(o.Ship.Country, "US") {
		tax, err := Tax(strings.ToUpper(o.Ship.Region), after)
		if err != nil {
			return Quote{}, err
		}
		q.Tax = tax
	}
	ship, err := Shipping("ground", Zone(strings.ToUpper(o.Ship.Country)), o.Items()*gramsPerItem)
	if err != nil {
		return Quote{}, err
	}
	q.Shipping = ship
	q.Total = after + q.Tax + q.Shipping
	return q, nil
}

// Format renders a quote total for display.
func (q Quote) Format() string {
	return money.Format(q.Total)
}
