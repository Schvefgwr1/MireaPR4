CREATE TABLE IF NOT EXISTS Products (
    product_id INT PRIMARY KEY generated always as identity,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES Categories(category_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);