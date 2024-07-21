-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id integer primary key,
    user_id text unique not null,
    username text not null,
    email text unique not null,
    user_password text not null,
    synopsis text,
    avatar text not null,
    background_image text not null,
    created_at datetime  not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
-- +goose StatementEnd