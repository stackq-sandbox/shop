// Package inventory tracks stock per warehouse and reserves it for orders.
package inventory

// SKU identifies a product.
type SKU string

// Item is a product in the catalog.
type Item struct {
	SKU     SKU
	Name    string
	WeightG int
	// Reorder is the available quantity at or below which stock is reordered.
	Reorder int
}

var catalog = []Item{
	{
		SKU:     "sku-001",
		Name:    "Widget 01",
		WeightG: 157,
		Reorder: 8,
	},
	{
		SKU:     "sku-002",
		Name:    "Gadget 02",
		WeightG: 194,
		Reorder: 11,
	},
	{
		SKU:     "sku-003",
		Name:    "Gizmo 03",
		WeightG: 231,
		Reorder: 14,
	},
	{
		SKU:     "sku-004",
		Name:    "Doohickey 04",
		WeightG: 268,
		Reorder: 17,
	},
	{
		SKU:     "sku-005",
		Name:    "Sprocket 05",
		WeightG: 305,
		Reorder: 20,
	},
	{
		SKU:     "sku-006",
		Name:    "Widget 06",
		WeightG: 342,
		Reorder: 23,
	},
	{
		SKU:     "sku-007",
		Name:    "Gadget 07",
		WeightG: 379,
		Reorder: 6,
	},
	{
		SKU:     "sku-008",
		Name:    "Gizmo 08",
		WeightG: 416,
		Reorder: 9,
	},
	{
		SKU:     "sku-009",
		Name:    "Doohickey 09",
		WeightG: 453,
		Reorder: 12,
	},
	{
		SKU:     "sku-010",
		Name:    "Sprocket 10",
		WeightG: 490,
		Reorder: 15,
	},
	{
		SKU:     "sku-011",
		Name:    "Widget 11",
		WeightG: 527,
		Reorder: 18,
	},
	{
		SKU:     "sku-012",
		Name:    "Gadget 12",
		WeightG: 564,
		Reorder: 21,
	},
	{
		SKU:     "sku-013",
		Name:    "Gizmo 13",
		WeightG: 601,
		Reorder: 24,
	},
	{
		SKU:     "sku-014",
		Name:    "Doohickey 14",
		WeightG: 638,
		Reorder: 7,
	},
	{
		SKU:     "sku-015",
		Name:    "Sprocket 15",
		WeightG: 675,
		Reorder: 10,
	},
	{
		SKU:     "sku-016",
		Name:    "Widget 16",
		WeightG: 712,
		Reorder: 13,
	},
	{
		SKU:     "sku-017",
		Name:    "Gadget 17",
		WeightG: 749,
		Reorder: 16,
	},
	{
		SKU:     "sku-018",
		Name:    "Gizmo 18",
		WeightG: 786,
		Reorder: 19,
	},
	{
		SKU:     "sku-019",
		Name:    "Doohickey 19",
		WeightG: 823,
		Reorder: 22,
	},
	{
		SKU:     "sku-020",
		Name:    "Sprocket 20",
		WeightG: 860,
		Reorder: 5,
	},
	{
		SKU:     "sku-021",
		Name:    "Widget 21",
		WeightG: 897,
		Reorder: 8,
	},
	{
		SKU:     "sku-022",
		Name:    "Gadget 22",
		WeightG: 934,
		Reorder: 11,
	},
	{
		SKU:     "sku-023",
		Name:    "Gizmo 23",
		WeightG: 971,
		Reorder: 14,
	},
	{
		SKU:     "sku-024",
		Name:    "Doohickey 24",
		WeightG: 1008,
		Reorder: 17,
	},
	{
		SKU:     "sku-025",
		Name:    "Sprocket 25",
		WeightG: 145,
		Reorder: 20,
	},
	{
		SKU:     "sku-026",
		Name:    "Widget 26",
		WeightG: 182,
		Reorder: 23,
	},
	{
		SKU:     "sku-027",
		Name:    "Gadget 27",
		WeightG: 219,
		Reorder: 6,
	},
	{
		SKU:     "sku-028",
		Name:    "Gizmo 28",
		WeightG: 256,
		Reorder: 9,
	},
	{
		SKU:     "sku-029",
		Name:    "Doohickey 29",
		WeightG: 293,
		Reorder: 12,
	},
	{
		SKU:     "sku-030",
		Name:    "Sprocket 30",
		WeightG: 330,
		Reorder: 15,
	},
	{
		SKU:     "sku-031",
		Name:    "Widget 31",
		WeightG: 367,
		Reorder: 18,
	},
	{
		SKU:     "sku-032",
		Name:    "Gadget 32",
		WeightG: 404,
		Reorder: 21,
	},
	{
		SKU:     "sku-033",
		Name:    "Gizmo 33",
		WeightG: 441,
		Reorder: 24,
	},
	{
		SKU:     "sku-034",
		Name:    "Doohickey 34",
		WeightG: 478,
		Reorder: 7,
	},
	{
		SKU:     "sku-035",
		Name:    "Sprocket 35",
		WeightG: 515,
		Reorder: 10,
	},
	{
		SKU:     "sku-036",
		Name:    "Widget 36",
		WeightG: 552,
		Reorder: 13,
	},
	{
		SKU:     "sku-037",
		Name:    "Gadget 37",
		WeightG: 589,
		Reorder: 16,
	},
	{
		SKU:     "sku-038",
		Name:    "Gizmo 38",
		WeightG: 626,
		Reorder: 19,
	},
	{
		SKU:     "sku-039",
		Name:    "Doohickey 39",
		WeightG: 663,
		Reorder: 22,
	},
	{
		SKU:     "sku-040",
		Name:    "Sprocket 40",
		WeightG: 700,
		Reorder: 5,
	},
}

var bySKU = func() map[SKU]Item {
	m := make(map[SKU]Item, len(catalog))
	for _, it := range catalog {
		m[it.SKU] = it
	}
	return m
}()

// Catalog lists every product.
func Catalog() []Item { return append([]Item(nil), catalog...) }

// Lookup returns the product with the given SKU.
func Lookup(sku SKU) (Item, bool) {
	it, ok := bySKU[sku]
	return it, ok
}
