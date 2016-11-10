CREATE TABLE IF NOT EXISTS project (
	id serial,
	name VARCHAR(255) NOT NULL,
	description text,
	location text,
	species text,
	pType text,
	visibility INTEGER NOT NULL,
	PRIMARY KEY(id)
);
