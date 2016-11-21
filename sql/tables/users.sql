CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(32),
	email text,
	name text,
	img_id INTEGER,
	location text DEFAULT 'Nowhere',
	season text DEFAULT 'All',
	growzone VARCHAR(32) DEFAULT 'undisclosed',
	specialty text DEFAULT 'None',
	about text DEFAULT '. . .' NOT NULL,
	hash bytea,
	salt bytea,
	ALGORITHM VARCHAR(10),
	PRIMARY KEY(id)
);
