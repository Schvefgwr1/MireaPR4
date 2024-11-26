CREATE TABLE IF NOT EXISTS Orders (
    order_id INT PRIMARY KEY generated always as identity,
    user_id INT,
    status_id INT,
    total_price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY (status_id) REFERENCES order_statuses(status_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);