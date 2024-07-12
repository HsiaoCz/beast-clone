-- +goose Up
CREATE TABLE IF NOT EXISTS sessionbs;
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS sessionbs;
-- +goose StatementBegin
-- +goose StatementEnd