// Package flags is a tiny feature-flag lookup. Production values live in
// config/env/prod/*.yaml; a SHOP_FLAG_<NAME> environment variable overrides.
package flags

import (
	"os"
	"strings"
)

// Flag names.
const (
	LegacyBanner = "legacy_banner"
	NewCheckout  = "new_checkout"
	Inventory    = "inventory"
	Reviews      = "reviews"
	Coupons      = "coupons"
	Loyalty      = "loyalty"
	Wishlist     = "wishlist"
	GiftNote     = "giftnote"
)

var defaults = map[string]bool{
	LegacyBanner: false,
	NewCheckout:  false,
	Inventory:    false,
	Reviews:      false,
	Coupons:      false,
	Loyalty:      false,
	Wishlist:     false,
	GiftNote:     false,
}

// Enabled reports whether the named flag is on.
func Enabled(name string) bool {
	if v, ok := os.LookupEnv("SHOP_FLAG_" + strings.ToUpper(name)); ok {
		return v == "1" || v == "true"
	}
	return defaults[name]
}
