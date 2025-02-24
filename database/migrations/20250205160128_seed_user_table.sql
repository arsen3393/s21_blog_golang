-- +goose Up
-- +goose StatementBegin
COPY users (name, email, password)
FROM '/migrations/admin_credentials.txt'
WITH (FORMAT csv, DELIMITER ',', HEADER false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE name = 'maria'
-- +goose StatementEnd
