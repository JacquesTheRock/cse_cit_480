INSERT INTO users(id,email,name,hash,salt, ALGORITHM) Values
('admin','noreply@bloomgenetics.tech','Admin','password','','PLAIN'),
('george','noreply@bloomgenetics.tech','Georgino','password','','PLAIN'),
('billy','noreply@bloomgenetics.tech','Billy','password','','PLAIN'),
('laura','noreply@bloomgenetics.tech','Laura Li','password','','PLAIN'),
('glenda','noreply@bloomgenetics.tech','Glenda Grinwald','password','','PLAIN'),
('guest','noreply@bloomgenetics.tech','Guest who','','','NOLOGIN');


INSERT INTO mail(src, dest, message,prev) VALUES
('admin','george','Hello George',NULL)
,('george','admin','Hello Admin',1);
