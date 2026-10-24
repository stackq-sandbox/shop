package orders

import (
	"github.com/stackq-sandbox/shop/internal/cart"
	"github.com/stackq-sandbox/shop/internal/util"
)

type Order struct {
	ID   string
	Cart cart.Cart
}

func (o Order) Summary() string {
	return o.ID + ": " + util.Format(o.Cart.Total())
}
