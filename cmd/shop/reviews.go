package main

import (
	"encoding/json"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/reviews"
)

type reviewRow struct {
	Name       string `json:"name"`
	Threshold  int64  `json:"threshold_cents"`
	PointsPer  int    `json:"points_per_dollar"`
	FreeReturn bool   `json:"free_return"`
}

// reviewsHandler lists the reviews tiers.
func reviewsHandler(w http.ResponseWriter, r *http.Request) {
	var rows []reviewRow
	for _, t := range reviews.Tiers() {
		rows = append(rows, reviewRow{t.Tier.String(), int64(t.Threshold), t.PointsPer, t.FreeReturn})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rows)
}
