// Package flags is a tiny feature-flag lookup. Production values live in
// config/env/prod/*.yaml; a SHOP_FLAG_<NAME> environment variable overrides.
package flags

import (
	"os"
	"strings"
)

var defaults = map[string]bool{
	"legacy_banner": false,
}

// Enabled reports whether the named flag is on.
func Enabled(name string) bool {
	if v, ok := os.LookupEnv("SHOP_FLAG_" + strings.ToUpper(name)); ok {
		return v == "1" || v == "true"
	}
	return defaults[name]
}
