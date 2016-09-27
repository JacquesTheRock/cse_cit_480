CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(32),
	email text,
	name text,
	location text,
	hash VARCHAR(255),
	salt VARCHAR(255),
	ALGORITHM VARCHAR(10),
	PRIMARY KEY(id)
)
