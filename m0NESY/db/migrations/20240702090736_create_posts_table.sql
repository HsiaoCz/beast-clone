-- +goose Up
CREATE TABLE IF NOT EXISTS posts(
    id integer primary key,
    user_id text unique not null,
    post_id text unique not null,
    content text not null,
    post_path text,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS posts;
-- +goose StatementBegin
-- +goose StatementEnd