BEGIN;

ALTER TABLE forms.form ADD COLUMN shareable_id VARCHAR(255) UNIQUE NOT NULL;
CREATE UNIQUE INDEX sharable_id_idx ON forms.form (shareable_id);
 
COMMIT;
