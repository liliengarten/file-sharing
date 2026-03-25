-- +goose Up
	CREATE TABLE role (
	    id   SERIAL PRIMARY KEY,
	    name VARCHAR(50) NOT NULL UNIQUE
	);
-- +goose Down
