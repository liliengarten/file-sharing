-- +goose Up
	CREATE TABLE board_pins (
		user_id   INT NOT NULL REFERENCES users(id)  ON DELETE CASCADE,
		pin_id    INT NOT NULL REFERENCES pins(id)   ON DELETE CASCADE,
		board_id  INT REFERENCES boards(id) ON DELETE SET NULL,
		added_at  TIMESTAMP NOT NULL DEFAULT NOW(),
		PRIMARY KEY (user_id, pin_id)
	);
-- +goose Down
