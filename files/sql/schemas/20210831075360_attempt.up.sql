CREATE TYPE language_type AS ENUM('python', 'go', 'js');

CREATE TABLE IF NOT EXISTS attempt (
    id UUID PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
    member_id UUID NOT NULL REFERENCES member(id),
    problem_id UUID NOT NULL REFERENCES problem(id),
    language language_type NOT NULL,
    is_solved boolean NOT NULL,
    score smallint NOT NULL,
    code TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT now(),
    update_at TIMESTAMP NOT NULL DEFAULT now()
);