-- +migrate Up
CREATE TABLE
    customers (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        phone VARCHAR(255) NOT NULL,
        UNIQUE(email)
    );
-- +migrate Down
DROP TABLE customers;