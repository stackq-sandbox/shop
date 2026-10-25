package orders

import (
	"testing"

	"github.com/stackq-sandbox/shop/internal/inventory"
)

func TestSummary(t *testing.T) {
	o := Order{ID: "o1"}
	o.Cart.Add("a", 2, 1250)
	if got, want := o.Summary(), "o1: $25.00"; got != want {
		t.Fatalf("Summary() = %q, want %q", got, want)
	}
}

func TestReserveHoldsEveryLine(t *testing.T) {
	inv := inventory.NewStore()
	inv.Receive("east-1", "sku-001", 5)
	inv.Receive("east-1", "sku-002", 5)
	o := Order{ID: "o1"}
	o.Cart.Add("sku-001", 2, 1000)
	o.Cart.Add("sku-002", 1, 500)
	o.Cart.Add("sku-001", 1, 1000) // same SKU again: quantities add up
	r, err := o.Reserve(inv, "east-1")
	if err != nil {
		t.Fatal(err)
	}
	if r.Lines["sku-001"] != 3 || r.Lines["sku-002"] != 1 {
		t.Fatalf("reservation = %+v", r.Lines)
	}
	if got := inv.Available("east-1", "sku-001"); got != 2 {
		t.Fatalf("available after reserving 3 of 5 = %d", got)
	}
}

func TestReserveFailsWithoutStock(t *testing.T) {
	o := Order{ID: "o1"}
	o.Cart.Add("sku-001", 1, 1000)
	if _, err := o.Reserve(inventory.NewStore(), "east-1"); err == nil {
		t.Fatal("an order with no stock behind it must not reserve")
	}
}
