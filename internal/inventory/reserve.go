package inventory

import (
	"errors"
	"fmt"
	"sort"
)

// Reservation is stock held for an order.
type Reservation struct {
	ID        string
	Warehouse WarehouseID
	Lines     map[SKU]int
}

func sortedSKUs(lines map[SKU]int) []SKU {
	out := make([]SKU, 0, len(lines))
	for sku := range lines {
		out = append(out, sku)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// Reserve holds every line or none of them.
func (s *Store) Reserve(id string, w WarehouseID, lines map[SKU]int) (Reservation, error) {
	if len(lines) == 0 {
		return Reservation{}, errors.New("inventory: nothing to reserve")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sku := range sortedSKUs(lines) {
		n := lines[sku]
		if n <= 0 {
			return Reservation{}, fmt.Errorf("inventory: cannot reserve %d of %q", n, sku)
		}
		l, err := s.level(w, sku)
		if err != nil {
			return Reservation{}, err
		}
		if l.Available() < n {
			return Reservation{}, fmt.Errorf("inventory: only %d of %q available in %s, need %d", l.Available(), sku, w, n)
		}
	}
	held := make(map[SKU]int, len(lines))
	for sku, n := range lines {
		s.levels[w][sku].Reserved += n
		held[sku] = n
	}
	return Reservation{ID: id, Warehouse: w, Lines: held}, nil
}

// Release gives a reservation back.
func (s *Store) Release(r Reservation) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sku := range sortedSKUs(r.Lines) {
		if l := s.levels[r.Warehouse][sku]; l == nil || l.Reserved < r.Lines[sku] {
			return fmt.Errorf("inventory: reservation %q holds more %q than the store has reserved", r.ID, sku)
		}
	}
	for sku, n := range r.Lines {
		s.levels[r.Warehouse][sku].Reserved -= n
	}
	return nil
}

// Commit turns a reservation into a shipment: the stock leaves the warehouse.
func (s *Store) Commit(r Reservation) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sku := range sortedSKUs(r.Lines) {
		if l := s.levels[r.Warehouse][sku]; l == nil || l.Reserved < r.Lines[sku] || l.OnHand < r.Lines[sku] {
			return fmt.Errorf("inventory: cannot commit %d of %q for reservation %q", r.Lines[sku], sku, r.ID)
		}
	}
	for sku, n := range r.Lines {
		l := s.levels[r.Warehouse][sku]
		l.Reserved -= n
		l.OnHand -= n
	}
	return nil
}
