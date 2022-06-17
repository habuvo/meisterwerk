-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS events (
    id uuid PRIMARY KEY,
    title TEXT NOT NULL,
    start_time timestamp,
    end_time timestamp,
    address text DEFAULT 'none',
    status text DEFAULT 'done'
    );
INSERT INTO events (id, title, start_time, end_time, status) VALUES (gen_random_uuid(), 'first event','2022-05-19 10:30:00','2022-05-19 11:30:00','done') ON CONFLICT DO NOTHING;
INSERT INTO events (id, title, start_time, end_time, status) VALUES (gen_random_uuid(), 'second event','2022-06-19 10:30:00','2022-06-19 11:30:00', 'pending') ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
