CREATE TABLE IF NOT EXISTS product_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES product_categories(id) ON DELETE SET NULL,
    name VARCHAR(150) NOT NULL,
    sku VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    price NUMERIC(12,2) NOT NULL,
    promo_price NUMERIC(12,2),
    stock INTEGER DEFAULT 0,
    weight NUMERIC(10,2),
    image_url TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO product_categories (name, description) VALUES
('Elektronik', 'Kategori produk elektronik'),
('Pakaian', 'Kategori produk pakaian')
ON CONFLICT DO NOTHING;

INSERT INTO products (category_id, name, sku, description, price, promo_price, stock, weight, image_url) VALUES
(1, 'Smartphone XYZ', 'SKU12345', 'Smartphone terbaru dengan fitur canggih', 5000000, 4500000, 100, 150, 'https://example.com/image1.jpg'),
(2, 'Kaos Polos', 'SKU54321', 'Kaos polos bahan katun', 100000, NULL, 50, 200, 'https://example.com/image2.jpg')
ON CONFLICT DO NOTHING;
