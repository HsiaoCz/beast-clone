-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
    id integer primary key,
    comment_id text unique not null,
    user_id text unique not null,
    post_id text unique not null,
    parent_id text,
    content text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone 
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS comments;
-- +goose StatementBegin
-- +goose StatementEnd