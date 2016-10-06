INSERT INTO users Values
('admin','noreply@bloomgenetics.tech','Admin','','password','','PLAIN'),
('george','noreply@bloomgenetics.tech','Georgino','','georgies password','','PLAIN');

INSERT INTO mail(src, dest, message,prev) VALUES
('admin','george','Hello George',NULL)
,('george','admin','Hello Admin',1);
