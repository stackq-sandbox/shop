package flags

import "testing"

func TestRefundsDefaultsOff(t *testing.T) {
	if Enabled(Refunds) {
		t.Fatal("refunds must default to off")
	}
	t.Setenv("SHOP_FLAG_REFUNDS", "1")
	if !Enabled(Refunds) {
		t.Fatal("env override should turn the flag on")
	}
}
