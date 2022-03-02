CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE accounts(
  id          uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  owned_by    uuid NOT NULL,
  status      VARCHAR(25) DEFAULT 'disabled' NOT NULL,
  enabled_at  TIMESTAMP,
  balance     INT DEFAULT 0 NOT NULL
);

CREATE TABLE access_tokens(
  id          SERIAL PRIMARY KEY,
  account_id  uuid NOT NULL,
  token       TEXT,
  created_at  TIMESTAMP DEFAULT now() NOT NULL
);