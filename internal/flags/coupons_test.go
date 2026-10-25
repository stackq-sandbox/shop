package flags

import "testing"

func TestCouponsDefaultsOff(t *testing.T) {
	if Enabled(Coupons) {
		t.Fatal("coupons must default to off")
	}
	t.Setenv("SHOP_FLAG_COUPONS", "1")
	if !Enabled(Coupons) {
		t.Fatal("env override should turn the flag on")
	}
}
