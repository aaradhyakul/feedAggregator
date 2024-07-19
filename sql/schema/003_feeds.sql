-- +goose Up
create table feeds(
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
drop table feeds;