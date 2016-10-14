CREATE TABLE IF NOT EXISTS crosses (
	id serial,
	project_id INTEGER,
	name VARCHAR(255),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS cross_parent (
	cross_id INTEGER,
	specimen_id INTEGER,
	PRIMARY KEY(cross_id,specimen_id)
);
