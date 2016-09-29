CREATE TABLE IF NOT EXISTS trait_t (
	id serial,
	name VARCHAR(255),
	weight FLOAT,
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS trait (
	id INTEGER,
	project_id INTEGER,
	name VARCHAR(255),
	class INTEGER,
	PRIMARY KEY(id)
);
