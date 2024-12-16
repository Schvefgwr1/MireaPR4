CREATE TABLE IF NOT EXISTS roles_permissions_rel (
    role_id INT,
    permission_id INT,
    FOREIGN KEY (role_id) REFERENCES Roles(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES Permissions(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);