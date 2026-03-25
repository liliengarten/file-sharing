-- +goose Up
	CREATE TABLE liked_pins (
		user_id   INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		pin_id    INT NOT NULL REFERENCES pins(id)  ON DELETE CASCADE,
		liked_at  TIMESTAMP NOT NULL DEFAULT NOW(),
		PRIMARY KEY (user_id, pin_id)
	);
-- +goose Down
