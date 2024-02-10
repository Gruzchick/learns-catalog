-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products
-- +goose StatementEnd
