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
