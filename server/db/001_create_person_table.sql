-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS public.PERSON
(
    ID serial NOT NULL,
    NAME varchar NOT NULL,
	SURNAME varchar NOT NULL,
	PATRONYMIC varchar,
	AGE integer NOT NULL,
    GENDER varchar NOT NULL,
	NATION varchar NOT NULL,
	CREATED_AT TIMESTAMP WITH TIME ZONE default NOW() NOT NULL,
	UPDATED_AT TIMESTAMP WITH TIME ZONE default NOW(),
    PRIMARY KEY (ID)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON PERSON
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

---- create above / drop below ----

drop table person;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
