BEGIN;

CREATE TABLE users (
  id integer PRIMARY KEY,
  username varchar(256) UNIQUE NOT NULL,
  email varchar NOT NULL,
  is_verified boolean DEFAULT false,
  password varchar(128) NOT NULL,
  created_at timestamp,
  updated_at timestamp,
  is_deleted boolean DEFAULT false
);

INSERT INTO users (id, username, email, password) VALUES (1, 'johnsnow', 'johnsnow@gmail.com', 'johnsnowpasswordhere');
INSERT INTO users (id, username, email, password) VALUES (2, 'raj', 'raj@gmail.com', 'rajspasswordhere');

COMMIT;
