CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE  user

CREATE TABLE IF NOT EXISTS "user" (
                                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now()
    );
