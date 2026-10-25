-- Expand step for the new checkout. Nothing here changes existing tables, so
-- it is safe to deploy before the code that uses it and to roll back by
-- ignoring it. The contract step (dropping orders.total_cents) is a later PR.

CREATE TABLE orders_v2 (
    id BIGSERIAL PRIMARY KEY,
    legacy_order_id BIGINT REFERENCES orders (id),
    coupon TEXT,
    payment TEXT NOT NULL,
    subtotal_cents BIGINT NOT NULL,
    discount_cents BIGINT NOT NULL DEFAULT 0,
    tax_cents BIGINT NOT NULL,
    shipping_cents BIGINT NOT NULL,
    total_cents BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE order_lines_v2 (
    order_id BIGINT NOT NULL REFERENCES orders_v2 (id),
    sku TEXT NOT NULL,
    qty INT NOT NULL CHECK (qty > 0),
    unit_cents BIGINT NOT NULL,
    PRIMARY KEY (order_id, sku)
);

CREATE TABLE shipping_addresses (
    order_id BIGINT PRIMARY KEY REFERENCES orders_v2 (id),
    line1 TEXT NOT NULL,
    line2 TEXT,
    city TEXT NOT NULL,
    region TEXT,
    postal TEXT NOT NULL,
    country CHAR(2) NOT NULL
);

CREATE TABLE shipments (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders_v2 (id),
    carrier TEXT NOT NULL,
    zone SMALLINT NOT NULL CHECK (zone BETWEEN 1 AND 8),
    grams INT NOT NULL CHECK (grams >= 0),
    price_cents BIGINT NOT NULL,
    shipped_at TIMESTAMPTZ
);

CREATE TABLE charges (
    id TEXT PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders_v2 (id),
    provider TEXT NOT NULL,
    amount_cents BIGINT NOT NULL,
    state TEXT NOT NULL CHECK (state IN ('authorized', 'captured', 'refunded')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE refunds (
    id BIGSERIAL PRIMARY KEY,
    charge_id TEXT NOT NULL REFERENCES charges (id),
    amount_cents BIGINT NOT NULL CHECK (amount_cents > 0),
    reason TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE tax_rates (
    region CHAR(2) PRIMARY KEY,
    name TEXT NOT NULL,
    bps INT NOT NULL CHECK (bps >= 0)
);

CREATE TABLE coupons (
    code TEXT PRIMARY KEY,
    kind TEXT NOT NULL CHECK (kind IN ('flat', 'percent')),
    value_cents BIGINT,
    percent SMALLINT,
    expires_at TIMESTAMPTZ
);

CREATE INDEX orders_v2_legacy_idx ON orders_v2 (legacy_order_id);
CREATE INDEX shipments_order_idx ON shipments (order_id);
CREATE INDEX charges_order_idx ON charges (order_id);
CREATE INDEX refunds_charge_idx ON refunds (charge_id);

INSERT INTO tax_rates (region, name, bps) VALUES ('AL', 'Alabama', 400);
INSERT INTO tax_rates (region, name, bps) VALUES ('AK', 'Alaska', 0);
INSERT INTO tax_rates (region, name, bps) VALUES ('AZ', 'Arizona', 560);
INSERT INTO tax_rates (region, name, bps) VALUES ('AR', 'Arkansas', 650);
INSERT INTO tax_rates (region, name, bps) VALUES ('CA', 'California', 725);
INSERT INTO tax_rates (region, name, bps) VALUES ('CO', 'Colorado', 290);
INSERT INTO tax_rates (region, name, bps) VALUES ('CT', 'Connecticut', 635);
INSERT INTO tax_rates (region, name, bps) VALUES ('DE', 'Delaware', 0);
INSERT INTO tax_rates (region, name, bps) VALUES ('FL', 'Florida', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('GA', 'Georgia', 400);
INSERT INTO tax_rates (region, name, bps) VALUES ('HI', 'Hawaii', 400);
INSERT INTO tax_rates (region, name, bps) VALUES ('ID', 'Idaho', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('IL', 'Illinois', 625);
INSERT INTO tax_rates (region, name, bps) VALUES ('IN', 'Indiana', 700);
INSERT INTO tax_rates (region, name, bps) VALUES ('IA', 'Iowa', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('KS', 'Kansas', 650);
INSERT INTO tax_rates (region, name, bps) VALUES ('KY', 'Kentucky', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('LA', 'Louisiana', 445);
INSERT INTO tax_rates (region, name, bps) VALUES ('ME', 'Maine', 550);
INSERT INTO tax_rates (region, name, bps) VALUES ('MD', 'Maryland', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('MA', 'Massachusetts', 625);
INSERT INTO tax_rates (region, name, bps) VALUES ('MI', 'Michigan', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('MN', 'Minnesota', 688);
INSERT INTO tax_rates (region, name, bps) VALUES ('MS', 'Mississippi', 700);
INSERT INTO tax_rates (region, name, bps) VALUES ('MO', 'Missouri', 423);
INSERT INTO tax_rates (region, name, bps) VALUES ('MT', 'Montana', 0);
INSERT INTO tax_rates (region, name, bps) VALUES ('NE', 'Nebraska', 550);
INSERT INTO tax_rates (region, name, bps) VALUES ('NV', 'Nevada', 685);
INSERT INTO tax_rates (region, name, bps) VALUES ('NH', 'New Hampshire', 0);
INSERT INTO tax_rates (region, name, bps) VALUES ('NJ', 'New Jersey', 663);
INSERT INTO tax_rates (region, name, bps) VALUES ('NM', 'New Mexico', 513);
INSERT INTO tax_rates (region, name, bps) VALUES ('NY', 'New York', 400);
INSERT INTO tax_rates (region, name, bps) VALUES ('NC', 'North Carolina', 475);
INSERT INTO tax_rates (region, name, bps) VALUES ('ND', 'North Dakota', 500);
INSERT INTO tax_rates (region, name, bps) VALUES ('OH', 'Ohio', 575);
INSERT INTO tax_rates (region, name, bps) VALUES ('OK', 'Oklahoma', 450);
INSERT INTO tax_rates (region, name, bps) VALUES ('OR', 'Oregon', 0);
INSERT INTO tax_rates (region, name, bps) VALUES ('PA', 'Pennsylvania', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('RI', 'Rhode Island', 700);
INSERT INTO tax_rates (region, name, bps) VALUES ('SC', 'South Carolina', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('SD', 'South Dakota', 450);
INSERT INTO tax_rates (region, name, bps) VALUES ('TN', 'Tennessee', 700);
INSERT INTO tax_rates (region, name, bps) VALUES ('TX', 'Texas', 625);
INSERT INTO tax_rates (region, name, bps) VALUES ('UT', 'Utah', 485);
INSERT INTO tax_rates (region, name, bps) VALUES ('VT', 'Vermont', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('VA', 'Virginia', 430);
INSERT INTO tax_rates (region, name, bps) VALUES ('WA', 'Washington', 650);
INSERT INTO tax_rates (region, name, bps) VALUES ('WV', 'West Virginia', 600);
INSERT INTO tax_rates (region, name, bps) VALUES ('WI', 'Wisconsin', 500);
INSERT INTO tax_rates (region, name, bps) VALUES ('WY', 'Wyoming', 400);
