-- Удаление существующих таблиц (если они есть)
DROP TABLE IF EXISTS purchase_items;
DROP TABLE IF EXISTS purchases;
DROP TABLE IF EXISTS deliveries;
DROP TABLE IF EXISTS price_change; -- Добавлена сюда
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS manufacturers;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS stores;
DROP TABLE IF EXISTS customers;


-- Создание таблицы магазинов
CREATE TABLE stores (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    phone VARCHAR(20)
);

-- Создание таблицы производителей
CREATE TABLE manufacturers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Создание таблицы категорий
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    store_id INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (store_id) REFERENCES stores(id)
);

-- Создание таблицы товаров
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
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

-- Создание таблицы клиентов
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    address TEXT
);

-- Создание таблицы покупок
CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    customer_id INT,
    store_id INT,
    product_id INT,
    quantity INT,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- Создание таблицы доставок
CREATE TABLE deliveries (
    id SERIAL PRIMARY KEY,
    order_id INT,
    status VARCHAR(255) NOT NULL,
    tracking_number VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES purchases(id)
);

-- Создание таблицы позиций покупки
CREATE TABLE purchase_items (
    id SERIAL PRIMARY KEY,
    purchase_id INT,
    product_id INT,
    quantity INT,
    price DECIMAL(10, 2),
    FOREIGN KEY (purchase_id) REFERENCES purchases(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Создание таблицы изменения цен (price_change)
CREATE TABLE price_change (
    id SERIAL PRIMARY KEY,
    product_id INT,
    old_price DECIMAL(10, 2),
    new_price DECIMAL(10, 2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

