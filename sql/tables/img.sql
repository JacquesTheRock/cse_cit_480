CREATE TABLE IF NOT EXISTS img (
	id bigserial,
	image bytea,
	ftype text,
	fsize INTEGER,
	PRIMARY KEY(id)
);
