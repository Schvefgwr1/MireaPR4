CREATE TABLE IF NOT EXISTS roles_permissions_rel (
    role_id INT,
    permission_id INT,
    FOREIGN KEY (role_id) REFERENCES Roles(role_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES Permissions(permission_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);