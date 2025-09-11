CREATE TABLE products (
      id BIGSERIAL PRIMARY KEY,
      shop_id BIGINT,
      code VARCHAR(16),
      name VARCHAR(255) NOT NULL,
      description TEXT,
      price NUMERIC(18,2) NOT NULL,
      created_at TIMESTAMP DEFAULT NOW(),
      updated_at TIMESTAMP DEFAULT NOW(),
      CONSTRAINT uq_product_shop UNIQUE (shop_id, code)
);
