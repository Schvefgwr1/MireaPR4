CREATE TABLE IF NOT EXISTS Shipments (
    shipment_id INT PRIMARY KEY generated always as identity,
    order_id INT,
    shipment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_id INT,
    address_id INT,
    FOREIGN KEY (address_id) REFERENCES Addresses(address_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES shipment_statuses(status_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);