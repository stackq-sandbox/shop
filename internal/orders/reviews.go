package orders

import "github.com/stackq-sandbox/shop/internal/reviews"

// Earn credits the customer with the points this order is worth.
func (o Order) EarnReviews(l *reviews.Ledger, customer string) (reviews.Account, reviews.Entry, error) {
	return l.Earn(customer, o.ID, o.Cart.Total())
}
