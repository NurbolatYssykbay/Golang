CREATE TABLE IF NOT EXISTS sequrity (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    safety level integer NOT NULL,
    sequrity text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
    );
