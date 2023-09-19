INSERT INTO orders (order_id, customer_id, order_time, order_status) VALUES ('b8fca141-b47f-4e7f-b002-63708dd9adac', 'harrison', now(), 'ready');
INSERT INTO products_ordered (order_id, product_id, quantity) VALUES ('b8fca141-b47f-4e7f-b002-63708dd9adac', 'banana', 1);
INSERT INTO products_ordered (order_id, product_id, quantity) VALUES ('b8fca141-b47f-4e7f-b002-63708dd9adac', 'orange', 1);
INSERT INTO orders (order_id, customer_id, order_time, order_status) VALUES ('d629e0f6-ee33-46d1-b6bc-26dc03d932ad', 'spencer', now(), 'ready');
INSERT INTO products_ordered (order_id, product_id, quantity) VALUES ('d629e0f6-ee33-46d1-b6bc-26dc03d932ad', 'beans', 50);
INSERT INTO orders (order_id, customer_id, order_time, order_status) VALUES ('c95e83d7-a9ea-4709-a40f-a355b0fac6ed', 'spencer', now(), 'ready');
INSERT INTO products_ordered (order_id, product_id, quantity) VALUES ('c95e83d7-a9ea-4709-a40f-a355b0fac6ed', 'beans', 50);
INSERT INTO products_ordered (order_id, product_id, quantity) VALUES ('c95e83d7-a9ea-4709-a40f-a355b0fac6ed', 'carrots', 100);