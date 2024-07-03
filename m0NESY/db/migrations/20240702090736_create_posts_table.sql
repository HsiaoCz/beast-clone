-- +goose Up
CREATE TABLE IF NOT EXISTS posts(
    id integer primary key,
    userID text unique not null,
    content text not null,
    created_at timestamp with time zone not null,
    staticPath text
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS posts;
-- +goose StatementBegin
-- +goose StatementEnd