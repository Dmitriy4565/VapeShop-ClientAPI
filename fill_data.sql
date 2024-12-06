-- Вставка данных в таблицу stores
INSERT INTO stores (name, address, phone) VALUES
('Магазин 1', 'Адрес 1', '123-456-7890'),
('Магазин 2', 'Адрес 2', '987-654-3210');

-- Вставка данных в таблицу manufacturers
INSERT INTO manufacturers (name) VALUES
('Производитель A'),
('Производитель B'),
('Производитель C');

-- Вставка данных в таблицу categories
INSERT INTO categories (name, store_id) VALUES
('Категория 1', 1),
('Категория 2', 1),
('Категория 3', 2);

-- Вставка данных в таблицу products
INSERT INTO products (name, description, price, image_url, category_id, manufacturer_id, stock, vape_type, power, battery_capacity, tank_capacity, coil_resistance, material, color, is_new, is_featured) VALUES
('Продукт 1', 'Описание продукта 1', 29.99, 'url1.jpg', 1, 1, 10, 'Pod', 20, 1000, 2, 0.8, 'Сталь', 'Черный', TRUE, TRUE),
('Продукт 2', 'Описание продукта 2', 39.99, 'url2.jpg', 2, 2, 5, 'Мод', 80, 2500, 4, 1.2, 'Алюминий', 'Серебряный', FALSE, FALSE);

-- Вставка данных в таблицу customers
INSERT INTO customers (first_name, last_name, email, password, phone, address) VALUES
('Иван', 'Иванов', 'ivan@example.com', 'password123', '111-222-3333', 'Адрес Иванова'),
('Петр', 'Петров', 'petr@example.com', 'password456', '444-555-6666', 'Адрес Петрова');

-- Вставка данных в таблицу purchases
INSERT INTO purchases (customer_id, store_id, product_id, quantity) VALUES
(1, 1, 1, 100),  -- Пример: customer_id=1, store_id=1, product_id=1, quantity=2
(2, 2, 2, 400);  -- Пример: customer_id=2, store_id=2, product_id=2, quantity=1

-- Вставка данных в таблицу deliveries
INSERT INTO deliveries (order_id, status, tracking_number, created_at, updated_at) VALUES
(1, 'Доставляется', '1Z999AA10123456785', NOW(), NOW()),
(2, 'Завершен', '1Z999AA10123456786', NOW(), NOW());

-- Вставка данных в таблицу purchase_items
INSERT INTO purchase_items (purchase_id, product_id, quantity, price) VALUES
(1, 1, 2, 29.99),
(2, 2, 1, 39.99);

-- Вставка данных в таблицу price_change (пример)
INSERT INTO price_change (product_id, old_price, new_price, created_at, updated_at) VALUES
(1, 35.00, 29.99, NOW(), NOW());

