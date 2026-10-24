CREATE TABLE carts (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE cart_items (
    cart_id     BIGINT NOT NULL REFERENCES carts (id),
    sku         TEXT   NOT NULL,
    qty         INT    NOT NULL CHECK (qty > 0),
    price_cents BIGINT NOT NULL,
    PRIMARY KEY (cart_id, sku)
);

CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    cart_id     BIGINT NOT NULL REFERENCES carts (id),
    total_cents BIGINT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
