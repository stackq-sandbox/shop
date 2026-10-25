package orders

import "github.com/stackq-sandbox/shop/internal/loyalty"

// Earn credits the customer with the points this order is worth.
func (o Order) Earn(l *loyalty.Ledger, customer string) (loyalty.Account, loyalty.Entry, error) {
	return l.Earn(customer, o.ID, o.Cart.Total())
}
