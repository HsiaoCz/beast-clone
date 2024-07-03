-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
    id integer primary key,
    userID text not null,
    postID text not null,
    parentID text,
    content text not null,
    created_at timestamp with time zone not null
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS comments;
-- +goose StatementBegin
-- +goose StatementEnd