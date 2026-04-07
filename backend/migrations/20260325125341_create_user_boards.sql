-- +goose Up
	CREATE TABLE user_boards (
		user_id    INT NOT NULL REFERENCES users(id)  ON DELETE CASCADE,
		board_id   INT NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
        invited_by INT REFERENCES users(id)  ON DELETE SET NULL,
		joined_at TIMESTAMP NOT NULL DEFAULT NOW(),
		PRIMARY KEY (user_id, board_id)
	);
-- +goose Down
