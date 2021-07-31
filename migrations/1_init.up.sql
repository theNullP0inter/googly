CREATE TABLE IF NOT EXISTS accounts (
  id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
  username VARCHAR(320) UNIQUE,
  deleted_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL
            DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL
            DEFAULT CURRENT_TIMESTAMP
            ON UPDATE CURRENT_TIMESTAMP
);