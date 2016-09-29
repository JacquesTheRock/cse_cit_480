CREATE TABLE IF NOT EXISTS logins (
	id bigserial,
	user_id VARCHAR(32) NOT NULL,
	name VARCHAR(255),
	key bytea NOT NULL,
	valid_until timestamp,
	PRIMARY KEY(id)
);
