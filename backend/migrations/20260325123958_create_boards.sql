-- +goose Up
	CREATE TABLE boards (
	    id          SERIAL PRIMARY KEY,
	    name        VARCHAR(255) NOT NULL,
	    description TEXT,
	    private     BOOLEAN NOT NULL DEFAULT TRUE
	);
-- +goose Down
