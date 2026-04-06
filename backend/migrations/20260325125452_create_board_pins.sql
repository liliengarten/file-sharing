-- +goose Up
	CREATE TABLE board_pins (
		pin_id    INT REFERENCES pins(id)   ON DELETE CASCADE,
		board_id  INT REFERENCES boards(id) ON DELETE CASCADE,
		added_at  TIMESTAMP NOT NULL DEFAULT NOW(),
		PRIMARY KEY (pin_id, board_id)
	);
-- +goose Down
