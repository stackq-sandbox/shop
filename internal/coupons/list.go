// Package coupons turns order totals into points and tiers.
package coupons

import "github.com/stackq-sandbox/shop/internal/money"

// Tier is a coupons level.
type Tier int

const (
	Bronze Tier = iota + 1
	Silver
	Gold
	Platinum
)

func (t Tier) String() string {
	switch t {
	case Bronze:
		return "bronze"
	case Silver:
		return "silver"
	case Gold:
		return "gold"
	case Platinum:
		return "platinum"
	}
	return "unknown"
}

// TierInfo describes a tier.
type TierInfo struct {
	Tier       Tier
	Threshold  money.Cents // lifetime spend needed
	PointsPer  int         // points per whole dollar
	FreeReturn bool
}

var tiers = []TierInfo{
	{Bronze, 0, 1, false},
	{Silver, 50000, 2, false},
	{Gold, 200000, 3, true},
	{Platinum, 1000000, 5, true},
}

// Tiers lists every tier, lowest first.
func Tiers() []TierInfo { return append([]TierInfo(nil), tiers...) }

// TierFor returns the tier reached by a lifetime spend.
func TierFor(lifetime money.Cents) TierInfo {
	cur := tiers[0]
	for _, t := range tiers {
		if lifetime >= t.Threshold {
			cur = t
		}
	}
	return cur
}
