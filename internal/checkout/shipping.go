package checkout

import (
	"fmt"

	"github.com/stackq-sandbox/shop/internal/money"
)

type Carrier struct {
	Name        string
	Base        money.Cents
	PerKg       money.Cents
	TransitDays int
}

var carriers = map[string]Carrier{
	"ground": {
		Name:        "Ground",
		Base:        499,
		PerKg:       120,
		TransitDays: 5,
	},
	"express": {
		Name:        "Express",
		Base:        1299,
		PerKg:       240,
		TransitDays: 2,
	},
	"overnight": {
		Name:        "Overnight",
		Base:        2499,
		PerKg:       480,
		TransitDays: 1,
	},
	"economy": {
		Name:        "Economy",
		Base:        349,
		PerKg:       90,
		TransitDays: 8,
	},
	"freight": {
		Name:        "Freight",
		Base:        4999,
		PerKg:       60,
		TransitDays: 10,
	},
	"pickup": {
		Name:        "Store pickup",
		Base:        0,
		PerKg:       0,
		TransitDays: 3,
	},
}

// Zone returns the shipping zone for a country. Zone 1 is domestic.
func Zone(country string) int {
	switch country {
	case "US":
		return 1
	case "CA", "MX":
		return 2
	case "GB", "IE", "DE", "FR", "NL", "BE", "ES", "IT", "PT", "AT", "CH":
		return 4
	case "SE", "NO", "DK", "FI", "PL":
		return 5
	case "AU", "JP":
		return 7
	default:
		return 8
	}
}

// Shipping prices a parcel of the given weight with a named carrier. Each zone
// beyond the first adds 15% to the carrier's price.
func Shipping(carrier string, zone, grams int) (money.Cents, error) {
	c, ok := carriers[carrier]
	if !ok {
		return 0, fmt.Errorf("checkout: unknown carrier %q", carrier)
	}
	if zone < 1 || zone > 8 {
		return 0, fmt.Errorf("checkout: zone %d out of range", zone)
	}
	if grams < 0 {
		return 0, fmt.Errorf("checkout: negative weight %d", grams)
	}
	price := c.Base + c.PerKg*money.Cents(grams)/1000
	pct := money.Cents(100 + 15*(zone-1))
	return price * pct / 100, nil
}
