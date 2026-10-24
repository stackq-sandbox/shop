package flags

import "testing"

func TestEnabled(t *testing.T) {
	if Enabled(LegacyBanner) {
		t.Fatal("legacy_banner should default to off")
	}
	t.Setenv("SHOP_FLAG_LEGACY_BANNER", "true")
	if !Enabled(LegacyBanner) {
		t.Fatal("env override should turn the flag on")
	}
}

func TestNewCheckoutDefaultsOff(t *testing.T) {
	if Enabled(NewCheckout) {
		t.Fatal("new_checkout must default to off")
	}
}
