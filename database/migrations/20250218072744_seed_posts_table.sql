-- +goose Up
-- +goose StatementBegin
COPY posts (title,text)
FROM '/migrations/posts_credentials.txt'
WITH (FORMAT csv, DELIMITER ',', QUOTE '"',HEADER false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM posts WHERE title IN ('# Добро пожаловать в мой блог!', '🌍 Как сделать мир лучше?', '## 🚀 Технологии, которые меняют будущее', '### 💡 Советы по продуктивности');
-- +goose StatementEnd
