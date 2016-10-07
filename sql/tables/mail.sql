CREATE TABLE IF NOT EXISTS mail (
	id bigserial,
	src VARCHAR(32) NOT NULL,
	dest VARCHAR(32) NOT NULL,
	prev INTEGER,
	arrival timestamp WITH time zone DEFAULT (now() at time zone 'utc'),
	subject VARCHAR(255) NOT NULL DEFAULT 'NO SUBJECT',
	message text NOT NULL,
	PRIMARY KEY(id)
);
