-- 0001_create_users_table.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_serie INTEGER NOT NULL,
    passport_number INTEGER NOT NULL,
    surname VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    patronymic VARCHAR,
    address VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 0001_create_users_table.down.sql
DROP TABLE IF EXISTS users;
