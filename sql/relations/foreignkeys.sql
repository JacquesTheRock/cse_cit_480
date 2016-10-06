ALTER TABLE logins
	ADD CONSTRAINT user_fk
	FOREIGN KEY(user_id)
	REFERENCES users(id);

ALTER TABLE roles 
	ADD CONSTRAINT user_fk
	FOREIGN KEY(user_id)
	REFERENCES users(id);
ALTER TABLE roles
	ADD CONSTRAINT project_fk
	FOREIGN KEY(project_id)
	REFERENCES project(id);
ALTER TABLE roles
	ADD CONSTRAINT role_type_fk
	FOREIGN KEY(role_id)
	REFERENCES role_t(id);

ALTER TABLE crosses
	ADD CONSTRAINT project_fk
	FOREIGN KEY(project_id)
	REFERENCES project(id);

ALTER TABLE specimen
	ADD CONSTRAINT crosses_fk
	FOREIGN KEY(cross_id)
	REFERENCES crosses(id);

ALTER TABLE trait
	ADD CONSTRAINT project_fk
	FOREIGN KEY(project_id)
	REFERENCES project(id);
ALTER TABLE trait
	ADD CONSTRAINT	traitt_fk
	FOREIGN KEY(class)
	REFERENCES trait_t(id);

ALTER TABLE specimen_trait
	ADD CONSTRAINT trait_fk
	FOREIGN KEY(trait_id)
	REFERENCES trait(id);
ALTER TABLE specimen_trait
	ADD CONSTRAINT specimen_fk
	FOREIGN KEY(specimen_id)
	REFERENCES specimen(id);


ALTER TABLE mail
	ADD CONSTRAINT src_fk
	FOREIGN KEY(src)
	REFERENCES users(id);
ALTER TABLE mail
	ADD CONSTRAINT dest_fk
	FOREIGN KEY(dest)
	REFERENCES users(id);
ALTER TABLE mail
	ADD CONSTRAINT prev_fk
	FOREIGN KEY(prev)
	REFERENCES mail(id);
