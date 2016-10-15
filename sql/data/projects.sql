INSERT INTO project(name,description,visibility) VALUES
('Example', 'This is just an example project', 1)
,('Test', 'This is a project for testing', 1);

INSERT INTO trait_T(name,weight) VALUES 
('recessive',0)
,('dominant',1)
,('custom',0.5);

INSERT INTO trait(project_id,name,class) VALUES
(1,'Green',2)
,(1,'Blue',1)
,(2,'Tall',1)
,(2,'Short',2);

INSERT INTO crosses(project_id,name) VALUES 
(1,'Cross1')
,(1,'Cross2');

INSERT INTO specimen(cross_id) VALUES
(1)
,(1)
,(2);

INSERT INTO specimen_trait(trait_id,specimen_id) VALUES
(1,1)
,(3,1)
,(2,2)
,(4,2);

INSERT INTO cross_parent(cross_id,specimen_id) VALUES
(2,1)
,(2,2);

