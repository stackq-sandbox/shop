package main

import (
	"encoding/json"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/refunds"
)

type refundRow struct {
	Name       string `json:"name"`
	Threshold  int64  `json:"threshold_cents"`
	PointsPer  int    `json:"points_per_dollar"`
	FreeReturn bool   `json:"free_return"`
}

// refundsHandler lists the refunds tiers.
func refundsHandler(w http.ResponseWriter, r *http.Request) {
	var rows []refundRow
	for _, t := range refunds.Tiers() {
		rows = append(rows, refundRow{t.Tier.String(), int64(t.Threshold), t.PointsPer, t.FreeReturn})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rows)
}
