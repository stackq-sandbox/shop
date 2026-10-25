# Loyalty

Customers earn points on every order and move up tiers by lifetime spend.

| Tier | Lifetime spend | Points per dollar | Free returns |
|------|----------------|-------------------|--------------|
| Bronze | $0 | 1 | no |
| Silver | $500 | 2 | no |
| Gold | $2,000 | 3 | yes |
| Platinum | $10,000 | 5 | yes |

## Rules

- Points are earned at the rate of the tier held *before* the order.
- An order earns at most once (`UNIQUE (order_id, reason)`).
- A refund takes back what the order earned, never below zero.
- Redeeming needs enough points; it never goes negative.

## Rollout

The `loyalty` flag is off by default. The tables ship first (additive
migration), then the code, then `config/env/prod/app.yaml` turns it on.
