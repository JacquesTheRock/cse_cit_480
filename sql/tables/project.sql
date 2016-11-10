CREATE TABLE IF NOT EXISTS project (
	id serial,
	name VARCHAR(255) NOT NULL,
	description text,
	location text DEFAULT '',
	species text DEFAULT '',
	pType text DEFAULT '',
	visibility INTEGER NOT NULL,
	PRIMARY KEY(id)
);
