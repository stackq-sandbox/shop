package orders

import (
	"github.com/stackq-sandbox/shop/internal/cart"
	"github.com/stackq-sandbox/shop/internal/inventory"
	"github.com/stackq-sandbox/shop/internal/money"
)

type Order struct {
	ID   string
	Cart cart.Cart
}

func (o Order) Summary() string {
	return o.ID + ": " + money.Format(o.Cart.Total())
}

// Reserve holds the stock for every line of the order in one warehouse.
func (o Order) Reserve(inv *inventory.Store, w inventory.WarehouseID) (inventory.Reservation, error) {
	lines := map[inventory.SKU]int{}
	for _, it := range o.Cart.Items {
		lines[inventory.SKU(it.SKU)] += it.Qty
	}
	return inv.Reserve(o.ID, w, lines)
}
