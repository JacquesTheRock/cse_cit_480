CREATE TABLE IF NOT EXISTS trait_t (
	id serial,
	name VARCHAR(255) NOT NULL,
	weight FLOAT,
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS trait (
	id bigserial,
	project_id INTEGER,
	name VARCHAR(255) NOT NULL,
	class INTEGER,
	pool INTEGER DEFAULT 0,
	PRIMARY KEY(id)
);
