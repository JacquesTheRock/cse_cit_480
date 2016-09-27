CREATE TABLE IF NOT EXISTS role_t (
	id serial,
	name VARCHAR(32),
	description text,
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS roles (
	user_id VARCHAR(32),
	project_id INTEGER,
	role_id INTEGER,
	PRIMARY KEY(user_id,project_id)
);
