-- +goose Up
CREATE TABLE IF NOT EXISTS admins(
    id integer primary key,
    user_id text not null,
    username text not null,
    user_password text not null,
    avatar text not null,
    email text unique not null,
    created_at datetime with time zone not null,
    updated_at datetime with time zone not null,
    deleted_at datetime with time zone not null
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS admins;
-- +goose StatementBegin
-- +goose StatementEnd