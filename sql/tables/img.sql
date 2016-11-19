CREATE TABLE IF NOT EXISTS img (
	id bigserial,
	image text NOT NULL,
	ftype text NOT NULL,
	fsize INTEGER DEFAULT 0,
	PRIMARY KEY(id)
);
