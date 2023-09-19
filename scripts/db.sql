CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE orders (
    order_id uuid PRIMARY KEY,
    customer_id text,
    order_time timestamp,
    order_status text
);

CREATE TABLE products_ordered (
    order_id uuid REFERENCES orders(order_id),
    product_id text,
    quantity real,
    PRIMARY KEY (order_id, product_id)
);