-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
    id integer primary key,
    user_id text not null,
    post_id text not null,
    parent_id text,
    content text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS comments;
-- +goose StatementBegin
-- +goose StatementEnd