package checkout

import (
	"fmt"

	"github.com/stackq-sandbox/shop/internal/money"
)

// StateTax is a flat state-level sales tax. Rates are illustrative fixtures,
// not tax advice.
type StateTax struct {
	Name   string
	BPS    int
	Exempt []string
}

var stateTax = map[string]StateTax{
	"AL": {
		Name:   "Alabama",
		BPS:    400,
		Exempt: []string{"grocery"},
	},
	"AK": {
		Name:   "Alaska",
		BPS:    0,
		Exempt: []string{"grocery", "clothing"},
	},
	"AZ": {
		Name:   "Arizona",
		BPS:    560,
		Exempt: []string{"grocery"},
	},
	"AR": {
		Name:   "Arkansas",
		BPS:    650,
		Exempt: []string{"grocery"},
	},
	"CA": {
		Name:   "California",
		BPS:    725,
		Exempt: []string{"grocery"},
	},
	"CO": {
		Name:   "Colorado",
		BPS:    290,
		Exempt: []string{"grocery"},
	},
	"CT": {
		Name:   "Connecticut",
		BPS:    635,
		Exempt: []string{"grocery"},
	},
	"DE": {
		Name:   "Delaware",
		BPS:    0,
		Exempt: []string{"grocery", "clothing"},
	},
	"FL": {
		Name:   "Florida",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"GA": {
		Name:   "Georgia",
		BPS:    400,
		Exempt: []string{"grocery"},
	},
	"HI": {
		Name:   "Hawaii",
		BPS:    400,
		Exempt: []string{"grocery"},
	},
	"ID": {
		Name:   "Idaho",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"IL": {
		Name:   "Illinois",
		BPS:    625,
		Exempt: []string{"grocery"},
	},
	"IN": {
		Name:   "Indiana",
		BPS:    700,
		Exempt: []string{"grocery"},
	},
	"IA": {
		Name:   "Iowa",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"KS": {
		Name:   "Kansas",
		BPS:    650,
		Exempt: []string{"grocery"},
	},
	"KY": {
		Name:   "Kentucky",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"LA": {
		Name:   "Louisiana",
		BPS:    445,
		Exempt: []string{"grocery"},
	},
	"ME": {
		Name:   "Maine",
		BPS:    550,
		Exempt: []string{"grocery"},
	},
	"MD": {
		Name:   "Maryland",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"MA": {
		Name:   "Massachusetts",
		BPS:    625,
		Exempt: []string{"grocery"},
	},
	"MI": {
		Name:   "Michigan",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"MN": {
		Name:   "Minnesota",
		BPS:    688,
		Exempt: []string{"grocery"},
	},
	"MS": {
		Name:   "Mississippi",
		BPS:    700,
		Exempt: []string{"grocery"},
	},
	"MO": {
		Name:   "Missouri",
		BPS:    423,
		Exempt: []string{"grocery"},
	},
	"MT": {
		Name:   "Montana",
		BPS:    0,
		Exempt: []string{"grocery", "clothing"},
	},
	"NE": {
		Name:   "Nebraska",
		BPS:    550,
		Exempt: []string{"grocery"},
	},
	"NV": {
		Name:   "Nevada",
		BPS:    685,
		Exempt: []string{"grocery"},
	},
	"NH": {
		Name:   "New Hampshire",
		BPS:    0,
		Exempt: []string{"grocery", "clothing"},
	},
	"NJ": {
		Name:   "New Jersey",
		BPS:    663,
		Exempt: []string{"grocery"},
	},
	"NM": {
		Name:   "New Mexico",
		BPS:    513,
		Exempt: []string{"grocery"},
	},
	"NY": {
		Name:   "New York",
		BPS:    400,
		Exempt: []string{"grocery"},
	},
	"NC": {
		Name:   "North Carolina",
		BPS:    475,
		Exempt: []string{"grocery"},
	},
	"ND": {
		Name:   "North Dakota",
		BPS:    500,
		Exempt: []string{"grocery"},
	},
	"OH": {
		Name:   "Ohio",
		BPS:    575,
		Exempt: []string{"grocery"},
	},
	"OK": {
		Name:   "Oklahoma",
		BPS:    450,
		Exempt: []string{"grocery"},
	},
	"OR": {
		Name:   "Oregon",
		BPS:    0,
		Exempt: []string{"grocery", "clothing"},
	},
	"PA": {
		Name:   "Pennsylvania",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"RI": {
		Name:   "Rhode Island",
		BPS:    700,
		Exempt: []string{"grocery"},
	},
	"SC": {
		Name:   "South Carolina",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"SD": {
		Name:   "South Dakota",
		BPS:    450,
		Exempt: []string{"grocery"},
	},
	"TN": {
		Name:   "Tennessee",
		BPS:    700,
		Exempt: []string{"grocery"},
	},
	"TX": {
		Name:   "Texas",
		BPS:    625,
		Exempt: []string{"grocery"},
	},
	"UT": {
		Name:   "Utah",
		BPS:    485,
		Exempt: []string{"grocery"},
	},
	"VT": {
		Name:   "Vermont",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"VA": {
		Name:   "Virginia",
		BPS:    430,
		Exempt: []string{"grocery"},
	},
	"WA": {
		Name:   "Washington",
		BPS:    650,
		Exempt: []string{"grocery"},
	},
	"WV": {
		Name:   "West Virginia",
		BPS:    600,
		Exempt: []string{"grocery"},
	},
	"WI": {
		Name:   "Wisconsin",
		BPS:    500,
		Exempt: []string{"grocery"},
	},
	"WY": {
		Name:   "Wyoming",
		BPS:    400,
		Exempt: []string{"grocery"},
	},
}

// Tax returns the sales tax on taxable for a US state.
func Tax(region string, taxable money.Cents) (money.Cents, error) {
	t, ok := stateTax[region]
	if !ok {
		return 0, fmt.Errorf("checkout: unknown region %q", region)
	}
	return taxable * money.Cents(t.BPS) / 10000, nil
}
