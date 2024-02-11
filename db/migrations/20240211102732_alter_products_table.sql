-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN price           integer NOT NULL DEFAULT 0,
    ADD COLUMN is_gaming       boolean,
    ADD COLUMN connection_type smallint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
