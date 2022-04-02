CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS sentries(id uuid DEFAULT uuid_generate_v4() NOT NULL, geolocation_id uuid, alias VARCHAR(50), PRIMARY KEY (id));
