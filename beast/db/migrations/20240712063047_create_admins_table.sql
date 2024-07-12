-- +goose Up
CREATE TABLE IF NOT EXISTS admins;
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS admins;
-- +goose StatementBegin
-- +goose StatementEnd