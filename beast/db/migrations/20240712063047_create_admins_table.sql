-- +goose Up
CREATE TABLE IF NOT EXISTS admins(
    id integer primary key,
    user_id text not null,
    username text not null,
    user_password text not null,
    avatar text not null,
    email text unique not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS admins;
-- +goose StatementBegin
-- +goose StatementEnd