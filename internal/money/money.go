package money

import "fmt"

// Cents is an amount of money in the smallest currency unit.
type Cents int64

// Format renders c as a dollar amount.
func Format(c Cents) string {
	return fmt.Sprintf("$%d.%02d", c/100, c%100)
}
