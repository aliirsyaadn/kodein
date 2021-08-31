CREATE TABLE IF NOT EXISTS project (
    id UUID PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
    member_id UUID NOT NULL REFERENCES member(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    technology TEXT NOT NULL,
    url TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now()
);