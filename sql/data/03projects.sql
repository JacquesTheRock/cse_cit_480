INSERT INTO project(id,name,description,visibility) VALUES
(0,'Default','This project is the root of all things', 1);

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
,(1,'Cross2')
,(1,'Cross3')
,(1,'Cross4');

INSERT INTO specimen(cross_id) VALUES
(1)
,(1)
,(1)
,(2)
,(2)
,(2)
,(3)
,(3)
,(3);

INSERT INTO specimen_trait(trait_id,specimen_id) VALUES
(1,1)
,(3,1)
,(2,2)
,(4,2);

INSERT INTO cross_parent(cross_id,specimen_id) VALUES
(2,1)
,(3,4)
,(3,5)
,(4,2);

INSERT INTO roles (user_id,project_id,role_id) VALUES
('guest',0,1)
,('admin',0,4)
,('admin',1,5)
,('admin',2,4)
,('guest',1,3)
,('guest',2,3);
