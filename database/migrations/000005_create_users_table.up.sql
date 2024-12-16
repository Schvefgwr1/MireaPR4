CREATE TABLE IF NOT EXISTS Users (
    user_id INT PRIMARY KEY generated always as identity,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    role_id INT,
    status_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES Roles(role_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY (status_id) REFERENCES User_Statuses(status_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);