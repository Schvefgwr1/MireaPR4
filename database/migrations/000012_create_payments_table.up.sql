CREATE TABLE IF NOT EXISTS Payments (
    payment_id INT PRIMARY KEY generated always as identity,
    order_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_id INT,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES payment_statuses(status_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);