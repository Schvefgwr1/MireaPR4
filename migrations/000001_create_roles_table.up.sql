CREATE TABLE IF NOT EXISTS Roles (
    role_id INT PRIMARY KEY generated always as identity,
    role_name VARCHAR(50) NOT NULL
);