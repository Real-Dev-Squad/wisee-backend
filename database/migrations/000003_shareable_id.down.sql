BEGIN;

DROP INDEX IF EXISTS shareable_id_idx ;
ALTER TABLE forms.form DROP COLUMN shareable_id;

COMMIT;
