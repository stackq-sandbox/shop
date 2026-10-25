package inventory

import (
	"fmt"
	"sort"
)

// WarehouseID identifies a warehouse.
type WarehouseID string

// Warehouse is a place stock is held.
type Warehouse struct {
	ID       WarehouseID
	Name     string
	Region   string
	Capacity int
}

var warehouses = []Warehouse{
	{
		ID:       "east-1",
		Name:     "Newark",
		Region:   "US-E",
		Capacity: 90000,
	},
	{
		ID:       "east-2",
		Name:     "Atlanta",
		Region:   "US-E",
		Capacity: 60000,
	},
	{
		ID:       "west-1",
		Name:     "Reno",
		Region:   "US-W",
		Capacity: 80000,
	},
	{
		ID:       "west-2",
		Name:     "Tacoma",
		Region:   "US-W",
		Capacity: 45000,
	},
	{
		ID:       "central-1",
		Name:     "Dallas",
		Region:   "US-C",
		Capacity: 70000,
	},
	{
		ID:       "eu-1",
		Name:     "Rotterdam",
		Region:   "EU",
		Capacity: 85000,
	},
	{
		ID:       "eu-2",
		Name:     "Lyon",
		Region:   "EU",
		Capacity: 40000,
	},
	{
		ID:       "apac-1",
		Name:     "Singapore",
		Region:   "APAC",
		Capacity: 55000,
	},
}

// Warehouses lists every warehouse, ordered by ID.
func Warehouses() []Warehouse {
	out := append([]Warehouse(nil), warehouses...)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Nearest returns the largest warehouse in a region.
func Nearest(region string) (Warehouse, error) {
	var best Warehouse
	for _, wh := range warehouses {
		if wh.Region == region && wh.Capacity > best.Capacity {
			best = wh
		}
	}
	if best.ID == "" {
		return Warehouse{}, fmt.Errorf("inventory: no warehouse in region %q", region)
	}
	return best, nil
}
