package cart

import (
	"github.com/stackq-sandbox/shop/internal/legacydiscount"
	"github.com/stackq-sandbox/shop/internal/money"
)

type Item struct {
	SKU   string
	Qty   int
	Price money.Cents
}

type Cart struct {
	Items []Item
}

func (c *Cart) Add(sku string, qty int, price money.Cents) {
	for i := range c.Items {
		if c.Items[i].SKU == sku {
			c.Items[i].Qty += qty
			return
		}
	}
	c.Items = append(c.Items, Item{SKU: sku, Qty: qty, Price: price})
}

func (c *Cart) Count() int {
	n := 0
	for _, it := range c.Items {
		n += it.Qty
	}
	return n
}

func (c *Cart) Subtotal() money.Cents {
	var sum money.Cents
	for _, it := range c.Items {
		sum += money.Cents(it.Qty) * it.Price
	}
	return sum
}

func (c *Cart) Total() money.Cents {
	return legacydiscount.Apply(c.Subtotal(), c.Count())
}
