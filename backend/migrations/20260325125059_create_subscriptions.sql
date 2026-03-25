-- +goose Up
	CREATE TABLE subscriptions (
		subscriber_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		target_id     INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
		PRIMARY KEY (subscriber_id, target_id),
		CONSTRAINT no_self_subscribe CHECK (subscriber_id <> target_id)
	);
-- +goose Down
