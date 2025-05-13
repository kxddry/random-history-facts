CREATE EXTENSION IF NOT EXISTS pg_trgm;


CREATE TABLE facts (
                       id SERIAL PRIMARY KEY,
                       fact TEXT NOT NULL,
                       normalized_fact TEXT NOT NULL
);

CREATE INDEX facts_normalized_fact_trgm_idx
    ON facts USING gin (normalized_fact gin_trgm_ops);
