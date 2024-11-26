CREATE TABLE IF NOT EXISTS Employees (
    employee_id INT PRIMARY KEY generated always as identity,
    user_id INT,
    position VARCHAR(100),
    department VARCHAR(100),
    phone VARCHAR(15),
    email VARCHAR(100),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);