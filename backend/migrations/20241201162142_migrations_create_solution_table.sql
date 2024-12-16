-- +goose Up
-- +goose StatementBegin
CREATE TABLE "solution"
(
    "id"          SERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    "description" TEXT,
    "playbook"    VARCHAR(255),
    created_at    TIMESTAMP NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "solution"
-- +goose StatementEnd
