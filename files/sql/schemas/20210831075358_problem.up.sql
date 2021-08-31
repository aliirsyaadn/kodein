CREATE TYPE difficulty_type AS ENUM('easy', 'normal', 'hard', 'insane');

CREATE TABLE IF NOT EXISTS problem (
    id UUID PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
    difficulty difficulty_type NOT NULL,
    grader_code TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now()
);