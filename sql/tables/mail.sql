CREATE TABLE IF NOT EXISTS mail (
	id INTEGER,
	src VARCHAR(32) NOT NULL,
	dest VARCHAR(32) NOT NULL,
	prev INTEGER,
	arrival timestamp WITH time zone DEFAULT (now() at time zone 'utc'),
	message text NOT NULL,
	PRIMARY KEY(id)
);
