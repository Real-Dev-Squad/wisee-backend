BEGIN;

ALTER TABLE forms.form DROP COLUMN shareable_id;

COMMIT;
