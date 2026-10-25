package main

import (
	"encoding/json"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/inventory"
)

type warehouseRow struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	Capacity int    `json:"capacity"`
}

// warehousesHandler lists the warehouses.
func warehousesHandler(w http.ResponseWriter, r *http.Request) {
	var rows []warehouseRow
	for _, wh := range inventory.Warehouses() {
		rows = append(rows, warehouseRow{string(wh.ID), wh.Name, wh.Region, wh.Capacity})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rows)
}
