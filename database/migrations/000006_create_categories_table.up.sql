CREATE TABLE IF NOT EXISTS Categories (
    category_id INT PRIMARY KEY generated always as identity,
    name VARCHAR(100) NOT NULL
);