package flags

import "testing"

func TestLoyaltyDefaultsOff(t *testing.T) {
	if Enabled(Loyalty) {
		t.Fatal("loyalty must default to off")
	}
	t.Setenv("SHOP_FLAG_LOYALTY", "1")
	if !Enabled(Loyalty) {
		t.Fatal("env override should turn the flag on")
	}
}
