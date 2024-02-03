BEGIN;

CREATE TYPE form_type_enum AS ENUM ('TEXT', 'TITLE', 'INPUT', 'RADIO');

-- Create forms table
CREATE TABLE IF NOT EXISTS forms (
    id BIGSERIAL PRIMARY KEY,
    created_by_id BIGINT,
    "created_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
    "updated_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
    "type" form_type_enum[][],
    content TEXT[][]
);

-- Create form_responses table
CREATE TABLE IF NOT EXISTS form_responses (
    id BIGSERIAL PRIMARY KEY,
    response_by_id BIGINT,
    content TEXT[][],
    form_id BIGINT,
    FOREIGN KEY (response_by_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (form_id) REFERENCES forms(id) ON DELETE SET NULL,
    "created_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC'),
    "updated_at" timestamp DEFAULT (NOW() AT TIME ZONE 'UTC')
);

-- Add foreign key constraint for forms table
ALTER TABLE forms
ADD CONSTRAINT fk_forms_created_by
FOREIGN KEY (created_by_id) REFERENCES users(id) ON DELETE SET NULL;

-- Add foreign key constraint for form_responses table
ALTER TABLE form_responses
ADD CONSTRAINT fk_form_responses_form
FOREIGN KEY (form_id) REFERENCES forms(id) ON DELETE SET NULL;

-- Add foreign key constraint for form_responses table
ALTER TABLE form_responses
ADD CONSTRAINT fk_form_responses_response_by
FOREIGN KEY (response_by_id) REFERENCES users(id) ON DELETE SET NULL;

COMMIT;
