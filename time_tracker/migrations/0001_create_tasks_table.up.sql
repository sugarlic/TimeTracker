-- 0001_create_tasks_table.up.sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

INSERT INTO tasks (name, description)
VALUES
    ('Бездельник', 'Ничего не делание'),
    ('Уборка', 'Уборка территории предприятия'),
    ('Учеба', 'Прохождение курсов по программированию');
