CREATE TABLE IF NOT EXISTS logins (
	id bigserial,
	user_id VARCHAR(32) NOT NULL,
	key VARCHAR(255) NOT NULL,
	valid_until timestamp,
	PRIMARY KEY(id)
);
