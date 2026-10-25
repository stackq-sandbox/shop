-- Expand step for refunds. Additive only: nothing here changes an existing
-- table, so it can ship before the code that uses it.

CREATE TABLE refunds_accounts (
    customer_id    TEXT PRIMARY KEY,
    tier           TEXT NOT NULL DEFAULT 'bronze',
    points         BIGINT NOT NULL DEFAULT 0 CHECK (points >= 0),
    lifetime_cents BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE refunds_entries (
    id          TEXT PRIMARY KEY,
    customer_id TEXT NOT NULL REFERENCES refunds_accounts (customer_id),
    order_id    TEXT NOT NULL,
    delta       BIGINT NOT NULL,
    reason      TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (order_id, reason)
);

CREATE INDEX refunds_entries_customer ON refunds_entries (customer_id, created_at);
