CREATE TABLE tickets (
    id serial PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    token_id NUMERIC NOT NULL UNIQUE,
    event_id INTEGER NOT NULL REFERENCES events(id),
    owner_id INTEGER NOT NULL REFERENCES users(id),
    burned_at TIMESTAMPTZ
)
