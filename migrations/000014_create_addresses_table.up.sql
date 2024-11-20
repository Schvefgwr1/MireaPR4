CREATE TABLE IF NOT EXISTS Addresses (
    address_id INT PRIMARY KEY generated always as identity,
    city VARCHAR(50) NOT NULL,
    street VARCHAR(50) NOT NULL,
    house INT NOT NULL,
    index VARCHAR(8) NOT NULL,
    flat INT not null
);