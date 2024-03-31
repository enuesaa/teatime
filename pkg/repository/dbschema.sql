CREATE TABLE teas (
  id         INTEGER PRIMARY KEY,
  teapod     VARCHAR(255) NOT NULL,
  collection VARCHAR(255) NOT NULL,
  -- format: `<teapod>:<name>`
  rid        VARCHAR(255) NOT NULL UNIQUE,
  value      JSON NOT NULL,
  created    DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated    DATETIME DEFAULT CURRENT_TIMESTAMP
);
