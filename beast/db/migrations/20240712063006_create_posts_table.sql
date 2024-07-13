-- +goose Up
CREATE TABLE IF NOT EXISTS posts(
    id integer primary key,
    userID text unique not null,
    postID text unique not null,
    content text not null,
    postPath text,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS posts;
-- +goose StatementBegin
-- +goose StatementEnd