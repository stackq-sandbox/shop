package orders

import (
	"github.com/stackq-sandbox/shop/internal/cart"
	"github.com/stackq-sandbox/shop/internal/money"
)

type Order struct {
	ID   string
	Cart cart.Cart
}

func (o Order) Summary() string {
	return o.ID + ": " + money.Format(o.Cart.Total())
}
