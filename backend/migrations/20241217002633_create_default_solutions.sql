-- +goose Up
-- +goose StatementBegin
INSERT INTO solution (name, description, playbook, updated_at)
VALUES ('First', 'This is a description of the first solution.', '1.yml', now());

INSERT INTO solution (name, description, playbook, updated_at)
VALUES ('Second', 'This is a description of the second solution.', '2.yml', now());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM solution
WHERE name = 'First' AND description = 'This is a description of the first solution.' AND playbook = '1.yml';

DELETE FROM solution
WHERE name = 'Second' AND description = 'This is a description of the second solution.' AND playbook = '2.yml';
-- +goose StatementEnd
