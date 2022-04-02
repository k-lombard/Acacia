CREATE TABLE IF NOT EXISTS images(id uuid DEFAULT uuid_generate_v4() NOT NULL, sentry_id uuid, content bytea NOT NULL, timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (id), CONSTRAINT fk_sentry FOREIGN KEY (sentry_id) REFERENCES sentries(id));