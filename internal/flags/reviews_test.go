package flags

import "testing"

func TestReviewsDefaultsOff(t *testing.T) {
	if Enabled(Reviews) {
		t.Fatal("reviews must default to off")
	}
	t.Setenv("SHOP_FLAG_REVIEWS", "1")
	if !Enabled(Reviews) {
		t.Fatal("env override should turn the flag on")
	}
}
