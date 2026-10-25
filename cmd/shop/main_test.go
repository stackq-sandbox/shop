package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func get(path string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	routes().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestCartTotal(t *testing.T) {
	rec := get("/cart/total")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "$40.48" {
		t.Fatalf("body = %q, want $40.48", got)
	}
}

func TestCheckoutV2IsFlagged(t *testing.T) {
	if rec := get("/checkout/v2"); rec.Code != http.StatusNotFound {
		t.Fatalf("flag off: status = %d, want 404", rec.Code)
	}
	t.Setenv("SHOP_FLAG_NEW_CHECKOUT", "true")
	if rec := get("/checkout/v2"); rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("flag on: status = %d, want 405 for GET", rec.Code)
	}
}

func TestInventoryWarehousesIsFlagged(t *testing.T) {
	if rec := get("/inventory/warehouses"); rec.Code != http.StatusNotFound {
		t.Fatalf("flag off: status = %d, want 404", rec.Code)
	}
	t.Setenv("SHOP_FLAG_INVENTORY", "true")
	rec := get("/inventory/warehouses")
	if rec.Code != http.StatusOK {
		t.Fatalf("flag on: status = %d", rec.Code)
	}
	var rows []map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &rows); err != nil || len(rows) != 8 {
		t.Fatalf("want 8 warehouses, got %d (%v)", len(rows), err)
	}
}

func TestLoyaltyTiersIsFlagged(t *testing.T) {
	if rec := get("/loyalty/tiers"); rec.Code != http.StatusNotFound {
		t.Fatalf("flag off: status = %d, want 404", rec.Code)
	}
	t.Setenv("SHOP_FLAG_LOYALTY", "true")
	rec := get("/loyalty/tiers")
	if rec.Code != http.StatusOK {
		t.Fatalf("flag on: status = %d", rec.Code)
	}
	var rows []map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &rows); err != nil || len(rows) != 4 {
		t.Fatalf("rows = %v, err = %v", rows, err)
	}
}
