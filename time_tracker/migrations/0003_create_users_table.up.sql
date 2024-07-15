-- 0003_create_users_table.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_serie INTEGER NOT NULL,
    passport_number INTEGER NOT NULL,
    surname VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    address VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);