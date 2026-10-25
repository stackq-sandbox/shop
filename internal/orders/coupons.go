package orders

import "github.com/stackq-sandbox/shop/internal/coupons"

// Earn credits the customer with the points this order is worth.
func (o Order) EarnCoupons(l *coupons.Ledger, customer string) (coupons.Account, coupons.Entry, error) {
	return l.Earn(customer, o.ID, o.Cart.Total())
}
