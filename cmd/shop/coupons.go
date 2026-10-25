package main

import (
	"encoding/json"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/coupons"
)

type couponRow struct {
	Name       string `json:"name"`
	Threshold  int64  `json:"threshold_cents"`
	PointsPer  int    `json:"points_per_dollar"`
	FreeReturn bool   `json:"free_return"`
}

// couponsHandler lists the coupons tiers.
func couponsHandler(w http.ResponseWriter, r *http.Request) {
	var rows []couponRow
	for _, t := range coupons.Tiers() {
		rows = append(rows, couponRow{t.Tier.String(), int64(t.Threshold), t.PointsPer, t.FreeReturn})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rows)
}
