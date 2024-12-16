CREATE TABLE IF NOT EXISTS Permissions (
    permission_id INT PRIMARY KEY generated always as identity,
    permission_name VARCHAR(100) NOT NULL
);