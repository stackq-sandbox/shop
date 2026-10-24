package checkout

import "testing"

func TestValidateAddressPostal(t *testing.T) {
	cases := []struct {
		country string
		good    string
		bad     string
	}{
		{"US", "94105", "9410"},
		{"CA", "K1A 0B1", "11111"},
		{"GB", "SW1A 1AA", "12345"},
		{"DE", "10115", "1011"},
		{"FR", "75001", "7500"},
		{"ES", "28013", "2801"},
		{"IT", "00184", "0018"},
		{"NL", "1011 AB", "1011"},
		{"BE", "1000", "100"},
		{"CH", "8001", "80"},
		{"AT", "1010", "10"},
		{"SE", "114 55", "1145"},
		{"NO", "0150", "015"},
		{"DK", "1050", "10"},
		{"FI", "00100", "0010"},
		{"PL", "00-950", "00950"},
		{"PT", "1000-001", "1000"},
		{"IE", "D02 X285", "12345"},
		{"AU", "2000", "200"},
		{"JP", "100-0001", "1000001"},
	}
	for _, c := range cases {
		a := Address{Line1: "1 Main St", City: "Town", Country: c.country}
		a.Postal = c.good
		if err := ValidateAddress(a); err != nil {
			t.Errorf("%s %q should be valid: %v", c.country, c.good, err)
		}
		a.Postal = c.bad
		if err := ValidateAddress(a); err == nil {
			t.Errorf("%s %q should be invalid", c.country, c.bad)
		}
	}
}

func TestValidateAddressRequiredFields(t *testing.T) {
	if err := ValidateAddress(Address{City: "Town", Postal: "94105", Country: "US"}); err == nil {
		t.Error("missing line 1 should fail")
	}
	if err := ValidateAddress(Address{Line1: "1 Main St", Postal: "94105", Country: "US"}); err == nil {
		t.Error("missing city should fail")
	}
	if err := ValidateAddress(Address{Line1: "1 Main St", City: "Town", Postal: "x", Country: "ZZ"}); err == nil {
		t.Error("unsupported country should fail")
	}
}
