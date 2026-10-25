package inventory

import (
	"fmt"
	"sort"
	"sync"
)

// Level is the stock of one SKU in one warehouse.
type Level struct {
	OnHand   int
	Reserved int
}

// Available is what can still be promised to an order.
func (l Level) Available() int { return l.OnHand - l.Reserved }

// Store holds stock levels. It is safe for concurrent use.
type Store struct {
	mu     sync.Mutex
	levels map[WarehouseID]map[SKU]*Level
}

func NewStore() *Store {
	s := &Store{levels: map[WarehouseID]map[SKU]*Level{}}
	for _, wh := range Warehouses() {
		s.levels[wh.ID] = map[SKU]*Level{}
	}
	return s
}

// level returns the entry for a SKU in a warehouse, creating it empty. The
// caller holds s.mu.
func (s *Store) level(w WarehouseID, sku SKU) (*Level, error) {
	m, ok := s.levels[w]
	if !ok {
		return nil, fmt.Errorf("inventory: unknown warehouse %q", w)
	}
	if _, ok := Lookup(sku); !ok {
		return nil, fmt.Errorf("inventory: unknown SKU %q", sku)
	}
	l := m[sku]
	if l == nil {
		l = &Level{}
		m[sku] = l
	}
	return l, nil
}

// Receive adds delivered stock.
func (s *Store) Receive(w WarehouseID, sku SKU, n int) error {
	if n <= 0 {
		return fmt.Errorf("inventory: cannot receive %d of %q", n, sku)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	l, err := s.level(w, sku)
	if err != nil {
		return err
	}
	l.OnHand += n
	return nil
}

// Available reports how much of a SKU can still be promised.
func (s *Store) Available(w WarehouseID, sku SKU) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if l := s.levels[w][sku]; l != nil {
		return l.Available()
	}
	return 0
}

// NeedsReorder lists SKUs held in a warehouse whose available quantity is at or
// below their reorder level.
func (s *Store) NeedsReorder(w WarehouseID) []SKU {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []SKU
	for sku, l := range s.levels[w] {
		if it, ok := Lookup(sku); ok && l.Available() <= it.Reorder {
			out = append(out, sku)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}
