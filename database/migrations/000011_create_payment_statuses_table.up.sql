CREATE TABLE IF NOT EXISTS payment_statuses (
    status_id INT PRIMARY KEY generated always as identity,
    status_name VARCHAR(20) NOT NULL
);