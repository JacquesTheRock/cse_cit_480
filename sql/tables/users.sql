CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(32),
	email text,
	name text,
	img bytea,
	location text DEFAULT 'Nowhere',
	season text DEFAULT 'All',
	growzone VARCHAR(32),
	specialty text DEFAULT 'None',
	hash bytea,
	salt bytea,
	ALGORITHM VARCHAR(10),
	PRIMARY KEY(id)
);
