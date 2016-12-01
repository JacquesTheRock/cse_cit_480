INSERT INTO project(id,name,description,visibility) VALUES
(0,'Default','This project is the root of all things', 1);

INSERT INTO project(name,description,visibility,location,species,ptype) VALUES
('Example', 'This is just an example project', 1,'detroit','Corn','Targeted')
,('Test', 'This is a project for testing', 1,'Nowheres Ville','Dog','See what happens');


INSERT INTO trait_T(name,weight) VALUES 
('recessive',0)
,('dominant',1)
,('custom',0.5);

INSERT INTO trait(project_id,name,class,pool) VALUES
(1,'Red',2,1)
,(1,'Yellow',1,1)
,(1,'Tall',1,2)
,(1,'Short',2,2)
,(1,'Juicy',1,3)
,(1,'Dry',2,3)
,(1,'Sweet',2,2)
,(1,'Sour',1,2);

INSERT INTO crosses(project_id,name) VALUES 
(1,'Start')
,(1,'EC2')
,(1,'EC3')
,(1,'EC4')
,(1,'EC5')
,(1,'EC6')
,(1,'Stable1');

INSERT INTO specimen(cross_id) VALUES
(1)
,(1)
,(1)
,(2)
,(2)
,(2)
,(3)
,(3)
,(3)
,(4)
,(4)
,(4)
,(5)
,(5)
,(5)
,(6)
,(6)
,(6)
,(7)
,(7)
,(7);

INSERT INTO specimen_trait(trait_id,specimen_id) VALUES
(1,1)
,(3,1)
,(5,1)
,(7,1)
,(2,2)
,(4,2)
,(6,2)
,(8,2)
,(1,3)
,(2,3)
,(3,3)
,(4,3)
,(5,3)
,(6,3)
,(7,3)
,(8,3);


INSERT INTO cross_parent(cross_id,specimen_id) VALUES
(2,1)
,(2,2)
,(3,3)
,(4,4)
,(4,5)
,(5,7)
,(5,10)
,(6,13)
,(7,16)
;
INSERT INTO roles (user_id,project_id,role_id) VALUES
('guest',0,1)
,('admin',0,4)
,('admin',1,5)
,('admin',2,4)
,('george',1,4)
,('billy',1,3)
,('laura',1,3)
,('glenda',2,5)
,('guest',1,3)
,('guest',2,3);
