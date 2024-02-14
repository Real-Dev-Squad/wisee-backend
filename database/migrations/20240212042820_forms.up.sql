BEGIN;

CREATE SCHEMA forms;

CREATE TYPE forms.status_type AS ENUM ('DRAFT', 'PUBLISHED');

CREATE TABLE forms.form (
    id bigserial PRIMARY KEY,
    content JSONB NOT NULL,
    created_by_id INT NOT NULL,
    owner_id INT NOT NULL,
    "status" forms.status_type DEFAULT 'DRAFT',
    created_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamp DEFAULT NULL
);

CREATE TABLE forms.metadata (
    id bigserial PRIMARY KEY,
    form_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    accepting_responses BOOLEAN NOT NULL DEFAULT FALSE,
    allow_guest_responses BOOLEAN NOT NULL DEFAULT TRUE,
    allow_multiple_responses BOOLEAN NOT NULL DEFAULT FALSE,
    send_confirmation_email_to_responde BOOLEAN NOT NULL DEFAULT FALSE,
    send_submission_email_to_owner BOOLEAN NOT NULL DEFAULT FALSE,
    valid_till timestamp,
    created_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamp DEFAULT NULL
);

CREATE TABLE forms.responses (
    id bigserial PRIMARY KEY,
    response_by_id INT NOT NULL,
    content JSONB NOT NULL,
    form_id INT NOT NULL,
    created_at timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamp DEFAULT NULL
);

ALTER TABLE forms.form
ADD CONSTRAINT fk_forms_users_created_by
FOREIGN KEY (created_by_id) 
REFERENCES users(id);

ALTER TABLE forms.form
ADD CONSTRAINT fk_forms_users_owner
FOREIGN KEY (owner_id) 
REFERENCES users(id);

ALTER TABLE forms.metadata
ADD CONSTRAINT fk_form_metadata_form
FOREIGN KEY (form_id) 
REFERENCES forms.form(id);

ALTER TABLE forms.responses
ADD CONSTRAINT fk_form_responses_users
FOREIGN KEY (response_by_id) 
REFERENCES users(id);

ALTER TABLE forms.responses
ADD CONSTRAINT fk_form_responses_form
FOREIGN KEY (form_id) 
REFERENCES forms.form(id);

COMMIT;
