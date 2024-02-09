BEGIN;

CREATE TYPE form_status_type AS ENUM ('DRAFT', 'PUBLISHED');

CREATE TABLE forms (
    id SERIAL PRIMARY KEY,
    content JSONB NOT NULL,
    created_by_id INT NOT NULL,
    owner_id INT NOT NULL,
    "status" form_status_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE form_metas (
    id SERIAL PRIMARY KEY,
    form_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    accepting_responses BOOLEAN NOT NULL DEFAULT FALSE,
    allow_guest_responses BOOLEAN NOT NULL DEFAULT TRUE,
    allow_multiple_responses BOOLEAN NOT NULL DEFAULT FALSE,
    send_confirmation_email_to_respondee BOOLEAN NOT NULL DEFAULT FALSE,
    send_submission_email_to_owner BOOLEAN NOT NULL DEFAULT FALSE,
    valid_till TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE form_responses (
    id SERIAL PRIMARY KEY,
    response_by_id INT NOT NULL,
    content JSONB NOT NULL,
    form_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE forms
ADD CONSTRAINT fk_forms_users_created_by
FOREIGN KEY (created_by_id) 
REFERENCES users(id);

ALTER TABLE forms
ADD CONSTRAINT fk_forms_users_owner
FOREIGN KEY (owner_id) 
REFERENCES users(id);

ALTER TABLE form_metas
ADD CONSTRAINT fk_form_metas_form
FOREIGN KEY (form_id) 
REFERENCES form(id);

ALTER TABLE form_responses
ADD CONSTRAINT fk_form_responses_users
FOREIGN KEY (response_by_id) 
REFERENCES users(id);

ALTER TABLE form_responses
ADD CONSTRAINT fk_form_responses_form
FOREIGN KEY (form_id) 
REFERENCES form(id);

COMMIT;
