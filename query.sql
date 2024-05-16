CREATE EXTENSION "uuid-ossp";

CREATE TABLE mst_authors (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    fullname VARCHAR(50) NULL,
    email VARCHAR(100) NULL UNIQUE,
    passwords VARCHAR(100) NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(100) NULL
);

INSERT INTO mst_authors(fullname, email, passwords, role) 
VALUES('Al Tsaqif', 'altsaqifnugraha19@gmail.com', '12345', 'user');