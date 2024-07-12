-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
    id integer primary key,
    commentID text unique not null,
    userID text unique not null,
    postID text unique not null,
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