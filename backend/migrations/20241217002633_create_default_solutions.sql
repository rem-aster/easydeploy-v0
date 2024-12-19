-- +goose Up
-- +goose StatementBegin
INSERT INTO solution (name, description, playbook, updated_at)
VALUES ('Factorio', 'Dedicated multiplayer server for popular game Factorio, where players can collaborate to build and manage factories in a procedurally generated world. The server hosts the game, allowing players to join and play together', 'factorio.yml', now());

INSERT INTO solution (name, description, playbook, updated_at)
VALUES ('Monica', 'Open source personal relationship management system, that lets you document your life.', 'monica.yml', now());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM solution
WHERE name = 'Factorio' AND description = 'Dedicated multiplayer server for popular game Factorio, where players can collaborate to build and manage factories in a procedurally generated world. The server hosts the game, allowing players to join and play together' AND playbook = 'factorio.yml';

DELETE FROM solution
WHERE name = 'Monica' AND description = 'Open source personal relationship management system, that lets you document your life.' AND playbook = 'monica.yml';
-- +goose StatementEnd
