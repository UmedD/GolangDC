CREATE TABLE products (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       price INT,
                       stock INT
);

INSERT INTO products (name, price, stock) VALUES ('toy', 120, 5),
                                                 ('bottle', 90, 20),
                                                 ('monitor', 200, 10),
                                                 ('computer', 350, 2);

SELECT * FROM products WHERE price < 100;

SELECT * FROM products WHERE stock >= 10;

UPDATE products SET price = price + 5 WHERE price < 100;

DELETE FROM products WHERE name LIKE '%TestProduct';

SELECT * FROM products ORDER BY price DESC;