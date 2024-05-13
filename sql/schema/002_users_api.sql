-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex') --cast to byte array, hash, encode in hexadecimal
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key