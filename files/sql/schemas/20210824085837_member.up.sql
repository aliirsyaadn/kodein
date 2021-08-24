CREATE TABLE IF NOT EXISTS member (
    id UUID PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
    name VARCHAR(60) NOT NULL,
    username VARCHAR(30) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email VARCHAR(60) NOT NULL,
    twitter VARCHAR(100),
    github VARCHAR(100),
    linkedin VARCHAR(100),
    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now()
);