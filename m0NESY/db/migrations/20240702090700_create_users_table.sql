-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id integer primary key,
    userID text unique not null,
    username text not null,
    email text unique not null,
    password text not null,
    synopsis text,
    avatar text not null,
    phone text unique not null,
    background_image text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
-- +goose StatementEnd