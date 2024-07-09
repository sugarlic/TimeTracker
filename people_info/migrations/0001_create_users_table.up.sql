-- 0001_create_users_table.up.sql
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


INSERT INTO users(passport_serie, passport_number, surname, name, patronymic, address, created_at, updated_at)
VALUES
(1234, 567890, 'Захаров', 'Илья', 'Валерьевич', 'Казань ул. Баумана д. 1 кв. 5', NOW(), NOW())