// Package wishlist keeps the items a customer saved for later.
package wishlist

import "sort"

// List is one customer's saved SKUs.
type List struct {
	skus map[string]bool
}

// Add saves sku; adding it twice keeps one copy.
func (l *List) Add(sku string) {
	if l.skus == nil {
		l.skus = map[string]bool{}
	}
	l.skus[sku] = true
}

// Remove forgets sku.
func (l *List) Remove(sku string) { delete(l.skus, sku) }

// Items returns the saved SKUs in order.
func (l *List) Items() []string {
	out := make([]string, 0, len(l.skus))
	for s := range l.skus {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}
