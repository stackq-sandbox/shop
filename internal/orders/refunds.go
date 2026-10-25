package orders

import "github.com/stackq-sandbox/shop/internal/refunds"

// Earn credits the customer with the points this order is worth.
func (o Order) EarnRefunds(l *refunds.Ledger, customer string) (refunds.Account, refunds.Entry, error) {
	return l.Earn(customer, o.ID, o.Cart.Total())
}
