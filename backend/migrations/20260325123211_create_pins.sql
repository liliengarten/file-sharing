-- +goose Up
	CREATE TABLE pins (
	    id          SERIAL PRIMARY KEY,
	    image       VARCHAR(500) NOT NULL,
	    description TEXT
	);
-- +goose Down
