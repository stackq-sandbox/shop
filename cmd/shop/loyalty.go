package main

import (
	"encoding/json"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/loyalty"
)

type tierRow struct {
	Name       string `json:"name"`
	Threshold  int64  `json:"threshold_cents"`
	PointsPer  int    `json:"points_per_dollar"`
	FreeReturn bool   `json:"free_return"`
}

// tiersHandler lists the loyalty tiers.
func tiersHandler(w http.ResponseWriter, r *http.Request) {
	var rows []tierRow
	for _, t := range loyalty.Tiers() {
		rows = append(rows, tierRow{t.Tier.String(), int64(t.Threshold), t.PointsPer, t.FreeReturn})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rows)
}
