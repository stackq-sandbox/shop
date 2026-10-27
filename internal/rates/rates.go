// Package rates quotes shipping rates per carrier.
package rates

// Quote is one carrier's price in cents.
type Quote struct {
	Carrier string
	Cents   int
}

// For returns the quotes for a parcel of the given weight in grams.
func For(grams int) []Quote {
	base := 300 + grams/10
	return []Quote{{"post", base}, {"courier", base * 2}}
}
