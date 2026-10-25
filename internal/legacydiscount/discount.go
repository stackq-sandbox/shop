// Package legacydiscount is the volume discount used by the old checkout.
package legacydiscount

import "github.com/stackq-sandbox/shop/internal/util"

// Apply takes 10% off orders of three or more items.
func Apply(sub util.Cents, items int) util.Cents {
	if items >= 3 {
		return sub - sub/10
	}
	return sub
}
