-- Удаление существующих таблиц (если они есть)
DROP TABLE IF EXISTS purchase_items;
DROP TABLE IF EXISTS purchases;
DROP TABLE IF EXISTS price_change;
DROP TABLE IF EXISTS deliveries;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS manufacturers;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS stores;
DROP TABLE IF EXISTS customers;

-- Создание таблицы магазинов
CREATE TABLE stores (
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 address VARCHAR(255),
 phone VARCHAR(20)
);

-- Создание таблицы категорий
CREATE TABLE categories (
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 store_id INT,
 FOREIGN KEY (store_id) REFERENCES stores(id)
);

-- Создание таблицы производителей
CREATE TABLE manufacturers (
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL
);

-- Создание таблицы товаров
CREATE TABLE products (
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 description TEXT,
 price DECIMAL(10, 2) NOT NULL,
 image_url VARCHAR(255),
 category_id INT,
 manufacturer_id INT,
 stock INT,
 vape_type VARCHAR(255),
 power INT,
 battery_capacity INT,
 tank_capacity INT,
 coil_resistance DECIMAL(4, 2),
 material VARCHAR(255),
 color VARCHAR(255),
 is_new BOOLEAN,
 is_featured BOOLEAN,
 FOREIGN KEY (category_id) REFERENCES categories(id),
 FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(id)
);

-- Создание таблицы жидкостей (если нужно)
CREATE TABLE liquids (
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 description TEXT,
 price DECIMAL(10, 2) NOT NULL,
 image_url VARCHAR(255),
 brand_id INT,
 nicotine_strength DECIMAL(4, 2),
 flavor VARCHAR(255),
 volume INT,
 vg_pg_ratio VARCHAR(255),
 FOREIGN KEY (brand_id) REFERENCES manufacturers(id)
);

-- Создание таблицы аксессуаров (если нужно)
CREATE TABLE accessories (
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 description TEXT,
 price DECIMAL(10, 2) NOT NULL,
 image_url VARCHAR(255),
 category_id INT,
 FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Создание таблицы клиентов
CREATE TABLE customers (
 id INT PRIMARY KEY AUTO_INCREMENT,
 first_name VARCHAR(255) NOT NULL,
 last_name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL UNIQUE,
 password VARCHAR(255) NOT NULL,
 phone VARCHAR(20),
 address TEXT
);

-- Создание таблицы доставок
CREATE TABLE deliveries (
 id INT PRIMARY KEY AUTO_INCREMENT,
 order_id INT,
  status VARCHAR(255) NOT NULL,
 tracking_number VARCHAR(255),
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 FOREIGN KEY (order_id) REFERENCES purchases(id)
);

-- Создание таблицы покупок
CREATE TABLE purchases (
 id INT PRIMARY KEY AUTO_INCREMENT,
 customer_id INT,
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- Создание таблицы позиций покупки
CREATE TABLE purchase_items (
 id INT PRIMARY KEY AUTO_INCREMENT,
 purchase_id INT,
 product_id INT,
 quantity INT,
 price DECIMAL(10, 2),
 FOREIGN KEY (purchase_id) REFERENCES purchases(id),
 FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Создание таблицы изменения цен
CREATE TABLE price_change (
 id INT PRIMARY KEY AUTO_INCREMENT,
 product_id INT,
 old_price DECIMAL(10, 2) NOT NULL,
 new_price DECIMAL(10, 2) NOT NULL,
 changed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Внешние ключи для доставок
ALTER TABLE deliveries
ADD CONSTRAINT fk_deliveries_orders FOREIGN KEY (order_id) REFERENCES purchases(id);

-- Внешние ключи для заказов
ALTER TABLE purchases
ADD CONSTRAINT fk_purchases_customers FOREIGN KEY (customer_id) REFERENCES customers(id);

-- Внешние ключи для позиций покупки
ALTER TABLE purchase_items
ADD CONSTRAINT fk_purchase_items_purchases FOREIGN KEY (purchase_id) REFERENCES purchases(id),
ADD CONSTRAINT fk_purchase_items_products FOREIGN KEY (product_id) REFERENCES products(id);

-- Внешние ключи для изменения цен
ALTER TABLE price_change
ADD CONSTRAINT fk_price_change_products FOREIGN KEY (product_id) REFERENCES products(id);

