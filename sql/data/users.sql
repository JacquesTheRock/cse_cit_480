INSERT INTO users(id,email,name,location,hash,salt, ALGORITHM) Values
('admin','noreply@bloomgenetics.tech','Admin','','password','','PLAIN'),
('george','noreply@bloomgenetics.tech','Georgino','','georgies password','','PLAIN'),
('guest','noreply@bloomgenetics.tech','Guest who', '','','','NOLOGIN');


INSERT INTO mail(src, dest, message,prev) VALUES
('admin','george','Hello George',NULL)
,('george','admin','Hello Admin',1);
