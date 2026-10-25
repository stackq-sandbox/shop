-- Expand step for inventory. Additive only: nothing here changes an existing
-- table, so it can ship before the code that uses it.

CREATE TABLE warehouses (
    id       TEXT PRIMARY KEY,
    name     TEXT NOT NULL,
    region   TEXT NOT NULL,
    capacity INT  NOT NULL CHECK (capacity > 0)
);

CREATE TABLE skus (
    sku       TEXT PRIMARY KEY,
    name      TEXT NOT NULL,
    weight_g  INT  NOT NULL CHECK (weight_g >= 0),
    reorder   INT  NOT NULL CHECK (reorder >= 0)
);

CREATE TABLE stock_levels (
    warehouse_id TEXT NOT NULL REFERENCES warehouses (id),
    sku          TEXT NOT NULL REFERENCES skus (sku),
    on_hand      INT  NOT NULL DEFAULT 0 CHECK (on_hand >= 0),
    reserved     INT  NOT NULL DEFAULT 0 CHECK (reserved >= 0),
    PRIMARY KEY (warehouse_id, sku),
    CHECK (reserved <= on_hand)
);

CREATE TABLE reservations (
    id           TEXT PRIMARY KEY,
    order_id     BIGINT REFERENCES orders (id),
    warehouse_id TEXT NOT NULL REFERENCES warehouses (id),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    state        TEXT NOT NULL CHECK (state IN ('held', 'released', 'committed'))
);

CREATE TABLE reservation_lines (
    reservation_id TEXT NOT NULL REFERENCES reservations (id),
    sku            TEXT NOT NULL REFERENCES skus (sku),
    qty            INT  NOT NULL CHECK (qty > 0),
    PRIMARY KEY (reservation_id, sku)
);

CREATE INDEX reservations_order_idx ON reservations (order_id);
CREATE INDEX stock_levels_sku_idx ON stock_levels (sku);

INSERT INTO warehouses (id, name, region, capacity) VALUES ('east-1', 'Newark', 'US-E', 90000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('east-2', 'Atlanta', 'US-E', 60000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('west-1', 'Reno', 'US-W', 80000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('west-2', 'Tacoma', 'US-W', 45000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('central-1', 'Dallas', 'US-C', 70000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('eu-1', 'Rotterdam', 'EU', 85000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('eu-2', 'Lyon', 'EU', 40000);
INSERT INTO warehouses (id, name, region, capacity) VALUES ('apac-1', 'Singapore', 'APAC', 55000);
