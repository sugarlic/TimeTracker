-- 0002_create_user_tasks_table.up.sql
CREATE TABLE user_tasks (
    id SERIAL PRIMARY KEY,
    surname VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    address VARCHAR(100) NOT NULL,
    task_id INTEGER NOT NULL REFERENCES tasks(id),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    total_minutes INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
