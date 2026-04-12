-- user password = password123 (bcrypt hash)
INSERT INTO users (id, name, email, password)
VALUES (
    uuid_generate_v4(),
    'Test User',
    'test@example.com',
    '$2a$12$8CeB7CrCSeMv./FnFuFZzOqHaL0E7Yqa0EcWwW2wbt9JRZ0K2/YtO'
);

-- project
INSERT INTO projects (id, name, description, owner_id)
SELECT uuid_generate_v4(), 'Demo Project', 'Sample project', id FROM users LIMIT 1;

-- tasks
INSERT INTO tasks (id, title, status, priority, project_id)
SELECT uuid_generate_v4(), 'Task 1', 'todo', 'high', id FROM projects LIMIT 1;

INSERT INTO tasks (id, title, status, priority, project_id)
SELECT uuid_generate_v4(), 'Task 2', 'in_progress', 'medium', id FROM projects LIMIT 1;

INSERT INTO tasks (id, title, status, priority, project_id)
SELECT uuid_generate_v4(), 'Task 3', 'done', 'low', id FROM projects LIMIT 1;