-- +goose Up
-- +goose StatementBegin
CREATE TABLE "deploy"
(
    "id"          SERIAL PRIMARY KEY,
    "status"      SMALLINT,
    "solution_id" INTEGER,
    "name"        VARCHAR(255),
    "ssh_address" VARCHAR(255),
    "ssh_key"     BYTEA,
    "extra"       JSONB,
    created_at    TIMESTAMP NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP
);
ALTER TABLE "deploy"
    ADD FOREIGN KEY ("solution_id") REFERENCES "solution" ("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "deploy"
-- +goose StatementEnd
