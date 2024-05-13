-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE -- user id stored in feeds table, references id of users in users table
    -- on delete cascade: when user is deleted feeds from that user should be deleted automatically
);

-- +goose down
DROP TABLE feeds;