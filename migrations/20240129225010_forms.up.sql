BEGIN;

-- Create ENUM type for form types
CREATE TYPE FORM_TYPE_ENUM AS ENUM (
  'TEXT_SHORT',
  'TEXT_LONG',
  'TEXT_NUMBER',
  'TEXT_EMAIL',
  'SINGLE_SELECT',
  'MULTI_SELECT',
  'FILE_IMAGE',
  'FILE_PDF',
  'LINEAR_SCALE'
);

-- Create forms table
CREATE TABLE forms (
  id bigserial PRIMARY KEY,
  created_by_id int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title text NOT NULL,
  "send_email_on_submission" boolean DEFAULT false,
  "created_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  "updated_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
);

-- Create questions table
CREATE TABLE questions (
  id bigserial PRIMARY KEY,
  form_id int NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
  title text NOT NULL,
  "type" FORM_TYPE_ENUM NOT NULL,
  points int,
  is_required boolean DEFAULT false,
  is_partial_marking_enabled boolean DEFAULT false,
  "created_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  "updated_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
);

-- Create options table
CREATE TABLE options (
  id bigserial PRIMARY KEY,
  question_id int NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
  "value" text NOT NULL,
  "created_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
  "updated_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
);


COMMIT;
