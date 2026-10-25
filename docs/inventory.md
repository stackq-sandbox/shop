# Inventory

Stock is tracked per warehouse and reserved for orders. Everything is behind the
`inventory` flag, which is off by default and switched on for production by
`config/env/prod/app.yaml`.

## Reservations

`Reserve` holds every line of an order or none of them. A reservation is then
either released (the order was cancelled) or committed (it shipped and the stock
leaves the warehouse).

## Warehouses

| ID | Name | Region | Capacity |
| -- | ---- | ------ | -------- |
| east-1 | Newark | US-E | 90000 |
| east-2 | Atlanta | US-E | 60000 |
| west-1 | Reno | US-W | 80000 |
| west-2 | Tacoma | US-W | 45000 |
| central-1 | Dallas | US-C | 70000 |
| eu-1 | Rotterdam | EU | 85000 |
| eu-2 | Lyon | EU | 40000 |
| apac-1 | Singapore | APAC | 55000 |

`Nearest(region)` picks the largest warehouse in a region.

## Rollback

Turn `inventory` off. The tables are additive and can stay.
