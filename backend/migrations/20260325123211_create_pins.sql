-- +goose Up
	CREATE TABLE pins (
	    id          SERIAL PRIMARY KEY,
        owner_id    INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	    image       VARCHAR(500) NOT NULL,
	    description TEXT
	);
-- +goose Down
