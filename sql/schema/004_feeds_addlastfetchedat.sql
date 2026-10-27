-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;


-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetched_at;

-- CONNECTION STRING: "postgres://postgres:lollakas@localhost:5432/gator"
