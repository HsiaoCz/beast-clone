-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id integer primary key,
    username text unique not null,
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
