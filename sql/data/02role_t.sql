INSERT INTO role_t(id,name,description) VALUES
(1,'Guest','Anonomous')
,(2,'User','System User')
,(3,'Member','Project Member')
,(4,'Admin','Project Admin')
,(5,'Owner','Project Owner');

INSERT INTO perm(id,page,action) VALUES
(1,'users','GET'),
(2,'users','POST'),
(3,'users_uid','GET'),
(4,'users_uid','PUT'),
(5,'users_uid_projects','GET'),
(6,'users_uid_mail','GET'),
(7,'users_uid_mail','POST'),
(8,'users_uid_mail_mid','GET'),
(9,'users_uid_mail_mid','PUT'),
(10,'users_uid_mail_mid','DELETE'),
(11,'projects','GET'),
(12,'projects','POST'),
(13,'projects_pid','GET'),
(14,'projects_pid','PUT'),
(15,'projects_pid','DELETE'),
(16,'projects_pid_traits','GET'),
(17,'projects_pid_traits','POST'),
(18,'projects_pid_traits_tid','GET'),
(19,'projects_pid_traits_tid','PUT'),
(20,'projects_pid_traits_tid','DELETE'),
(21,'projects_pid_crosses','GET'),
(22,'projects_pid_crosses','POST'),
(23,'projects_pid_crosses_cid','GET'),
(24,'projects_pid_crosses_cid','PUT'),
(25,'projects_pid_crosses_cid','DELETE'),
(26,'projects_pid_crosses_cid_candidates','GET'),
(27,'projects_pid_crosses_cid_candidates','POST'),
(28,'projects_pid_crosses_cid_candidates_cnid','GET'),
(29,'projects_pid_crosses_cid_candidates_cnid','PUT'),
(30,'projects_pid_crosses_cid_candidates_cnid','DELETE'),
(31,'projects_pid_treview','GET'),
(32,'projects_pid_treview_cid','GET'),
(33,'auth','GET'),
(34,'auth','POST'),
(35,'auth','DELETE'),
(36,'projects_pid_roles','GET'),
(37,'projects_pid_roles','POST'),
(38,'projects_pid_roles_uid','GET'),
(39,'projects_pid_roles_uid','PUT'),
(40,'projects_pid_roles_uid','DELETE'),
(41,'images','POST'),
(42,'images_iid','GET'),
(43,'images_iid','DELETE'),
(44,'projects_pid_candidates','GET'),
(45,'projects_pid_candidates_cnid','GET'),
(46,'projects_pid_candidates_cnid','PUT')
;

INSERT INTO role_perm(role_id, perm_id) VALUES 
(1,33)
,(1,34)
,(1,35)
,(1,2)
,(1,42);

INSERT INTO role_perm(role_id,perm_id)
	SELECT 2, perm_id FROM role_perm WHERE role_id=1;

INSERT INTO role_perm(role_id,perm_id) VALUES
(2,1)
,(2,3)
,(2,4)
,(2,5)
,(2,6)
,(2,7)
,(2,8)
,(2,9)
,(2,10)
,(2,11)
,(2,12)
,(2,13)
,(2,41)
,(2,43);

INSERT INTO role_perm(role_id,perm_id)
	SELECT 3, perm_id FROM role_perm WHERE role_id=2;

INSERT INTO role_perm(role_id,perm_id) VALUES
(3,16)
,(3,18)
,(3,21)
,(3,23)
,(3,26)
,(3,27)
,(3,28)
,(3,31)
,(3,32)
,(3,36)
,(2,44)
,(2,45);


INSERT INTO role_perm(role_id,perm_id)
	SELECT 4, perm_id FROM role_perm WHERE role_id=3;

INSERT INTO role_perm(role_id,perm_id) VALUES
(4,14)
,(4,17)
,(4,19)
,(4,20)
,(4,22)
,(4,24)
,(4,25)
,(4,29)
,(4,30)
,(4,46);

INSERT INTO role_perm(role_id,perm_id)
	SELECT 5, perm_id FROM role_perm WHERE role_id=4;

INSERT INTO role_perm(role_id,perm_id) VALUES
(5,15),
(5,37),
(5,38),
(5,39),
(5,40);
