package flags

import "testing"

func TestEnabled(t *testing.T) {
	if Enabled("legacy_banner") {
		t.Fatal("legacy_banner should default to off")
	}
	t.Setenv("SHOP_FLAG_LEGACY_BANNER", "true")
	if !Enabled("legacy_banner") {
		t.Fatal("env override should turn the flag on")
	}
}
