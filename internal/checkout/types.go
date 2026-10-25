// Package checkout prices and places orders for the rewritten checkout.
package checkout

import "github.com/stackq-sandbox/shop/internal/money"

type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2,omitempty"`
	City    string `json:"city"`
	Region  string `json:"region"`
	Postal  string `json:"postal"`
	Country string `json:"country"`
}

type Line struct {
	SKU  string      `json:"sku"`
	Qty  int         `json:"qty"`
	Unit money.Cents `json:"unit_cents"`
}

type Order struct {
	ID      string  `json:"id"`
	Lines   []Line  `json:"lines"`
	Ship    Address `json:"ship"`
	Coupon  string  `json:"coupon,omitempty"`
	Payment string  `json:"payment"`
}

type Quote struct {
	Subtotal money.Cents `json:"subtotal_cents"`
	Discount money.Cents `json:"discount_cents"`
	Tax      money.Cents `json:"tax_cents"`
	Shipping money.Cents `json:"shipping_cents"`
	Total    money.Cents `json:"total_cents"`
}

func (o Order) Items() int {
	n := 0
	for _, l := range o.Lines {
		n += l.Qty
	}
	return n
}

func (o Order) Subtotal() money.Cents {
	var sum money.Cents
	for _, l := range o.Lines {
		sum += money.Cents(l.Qty) * l.Unit
	}
	return sum
}
