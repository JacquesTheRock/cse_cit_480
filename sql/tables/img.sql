CREATE TABLE IF NOT EXISTS img (
	id bigserial,
	image bytea NOT NULL,
	ftype text NOT NULL,
	fsize INTEGER DEFAULT 0,
	PRIMARY KEY(id)
);
