CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(32),
	email text,
	name text,
	location text,
	hash bytea,
	salt bytea,
	ALGORITHM VARCHAR(10),
	PRIMARY KEY(id)
);
