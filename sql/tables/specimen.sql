CREATE TABLE IF NOT EXISTS specimen (
	id serial,
	cross_id INTEGER,
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS specimen_trait(
	trait_id INTEGER,
	specimen_id INTEGER,
	PRIMARY KEY(trait_id,specimen_id)
);
