package inventory

import (
	"reflect"
	"testing"
)

func TestCatalogAndWarehouses(t *testing.T) {
	if len(Catalog()) != 40 {
		t.Fatalf("catalog has %d items", len(Catalog()))
	}
	if it, ok := Lookup("sku-001"); !ok || it.Name != "Widget 01" {
		t.Fatalf("Lookup(sku-001) = %+v, %v", it, ok)
	}
	if _, ok := Lookup("nope"); ok {
		t.Fatal("unknown SKU found")
	}
	ws := Warehouses()
	if len(ws) != 8 || ws[0].ID != "apac-1" || ws[7].ID != "west-2" {
		t.Fatalf("warehouses = %+v", ws)
	}
}

func TestNearestPicksTheLargestInTheRegion(t *testing.T) {
	cases := map[string]WarehouseID{"US-E": "east-1", "US-W": "west-1", "EU": "eu-1", "APAC": "apac-1", "US-C": "central-1"}
	for region, want := range cases {
		got, err := Nearest(region)
		if err != nil || got.ID != want {
			t.Errorf("Nearest(%q) = %v, %v; want %s", region, got.ID, err, want)
		}
	}
	if _, err := Nearest("MARS"); err == nil {
		t.Error("an unknown region should fail")
	}
}

func TestReceiveAndAvailable(t *testing.T) {
	s := NewStore()
	if err := s.Receive("east-1", "sku-001", 10); err != nil {
		t.Fatal(err)
	}
	if got := s.Available("east-1", "sku-001"); got != 10 {
		t.Fatalf("Available = %d", got)
	}
	for _, c := range []struct {
		w   WarehouseID
		sku SKU
		n   int
	}{{"mars", "sku-001", 1}, {"east-1", "nope", 1}, {"east-1", "sku-001", 0}} {
		if err := s.Receive(c.w, c.sku, c.n); err == nil {
			t.Errorf("Receive(%v) should fail", c)
		}
	}
}

func TestReserveIsAllOrNothing(t *testing.T) {
	s := NewStore()
	s.Receive("east-1", "sku-001", 5)
	s.Receive("east-1", "sku-002", 1)
	if _, err := s.Reserve("o1", "east-1", map[SKU]int{"sku-001": 2, "sku-002": 3}); err == nil {
		t.Fatal("reserving more than is available must fail")
	}
	if got := s.Available("east-1", "sku-001"); got != 5 {
		t.Fatalf("a failed reservation held stock: %d", got)
	}
	r, err := s.Reserve("o2", "east-1", map[SKU]int{"sku-001": 2, "sku-002": 1})
	if err != nil {
		t.Fatal(err)
	}
	if s.Available("east-1", "sku-001") != 3 || s.Available("east-1", "sku-002") != 0 {
		t.Fatal("reserved stock should not be available")
	}
	if err := s.Release(r); err != nil || s.Available("east-1", "sku-001") != 5 {
		t.Fatalf("Release: %v, available %d", err, s.Available("east-1", "sku-001"))
	}
}

func TestCommitRemovesStock(t *testing.T) {
	s := NewStore()
	s.Receive("west-1", "sku-003", 4)
	r, _ := s.Reserve("o1", "west-1", map[SKU]int{"sku-003": 3})
	if err := s.Commit(r); err != nil {
		t.Fatal(err)
	}
	if got := s.Available("west-1", "sku-003"); got != 1 {
		t.Fatalf("after committing 3 of 4, available = %d", got)
	}
	if err := s.Commit(r); err == nil {
		t.Fatal("committing the same reservation twice must fail")
	}
}

func TestNeedsReorder(t *testing.T) {
	s := NewStore()
	it, _ := Lookup("sku-001")
	s.Receive("eu-1", "sku-001", it.Reorder+10)
	s.Receive("eu-1", "sku-002", 1)
	if got := s.NeedsReorder("eu-1"); !reflect.DeepEqual(got, []SKU{"sku-002"}) {
		t.Fatalf("NeedsReorder = %v", got)
	}
}
