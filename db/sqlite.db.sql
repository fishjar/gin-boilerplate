BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `userrole` (
	`role_id`	blob NOT NULL,
	`user_id`	blob NOT NULL,
	PRIMARY KEY(`role_id`,`user_id`)
);
INSERT INTO `userrole` VALUES ('89d17436-f83b-4644-9ebe-feb5ad51e4d9','3cc8591c-d77f-4715-a48c-b2e0a44e04e4');
INSERT INTO `userrole` VALUES ('87934398-0bc9-4ffb-98be-18adeefe4967','3cc8591c-d77f-4715-a48c-b2e0a44e04e4');
INSERT INTO `userrole` VALUES ('87af54fd-18a8-488c-9438-52d7f032be31','3cc8591c-d77f-4715-a48c-b2e0a44e04e4');
INSERT INTO `userrole` VALUES ('87934398-0bc9-4ffb-98be-18adeefe4967','a22824b3-c6eb-4757-a93c-6bff97337824');
INSERT INTO `userrole` VALUES ('87af54fd-18a8-488c-9438-52d7f032be31','a22824b3-c6eb-4757-a93c-6bff97337824');
INSERT INTO `userrole` VALUES ('87af54fd-18a8-488c-9438-52d7f032be31','28ae1605-73b8-4f9c-b393-d38d2e1addc7');
CREATE TABLE IF NOT EXISTS `usergroup` (
	`id`	blob NOT NULL,
	`created_at`	datetime NOT NULL,
	`update_at`	datetime NOT NULL,
	`deleted_at`	datetime,
	`user_id`	blob NOT NULL,
	`group_id`	blob NOT NULL,
	`level`	TINYINT DEFAULT 0,
	`join_time`	DATETIME,
	PRIMARY KEY(`id`)
);
INSERT INTO `usergroup` VALUES ('c3b60aaf-cce0-42b0-a413-304cfd33abd6','2020-05-28 23:55:24.331874017+08:00','2020-05-28 23:55:24.331874017+08:00',NULL,'3cc8591c-d77f-4715-a48c-b2e0a44e04e4','6dd8d5e4-b430-4a87-97f5-7b7e5367a692',1,NULL);
INSERT INTO `usergroup` VALUES ('ca3490b5-6bc3-42f3-b0c9-ea6f2fde8bfb','2020-05-28 23:55:24.340133064+08:00','2020-05-28 23:55:24.340133064+08:00',NULL,'3cc8591c-d77f-4715-a48c-b2e0a44e04e4','f4af0d66-6eff-4420-b95e-b67c6c662739',1,NULL);
INSERT INTO `usergroup` VALUES ('bc13f1c0-318d-454f-bf34-3ba63ebf5ce1','2020-05-28 23:55:24.345831693+08:00','2020-05-28 23:55:24.345831693+08:00',NULL,'a22824b3-c6eb-4757-a93c-6bff97337824','6dd8d5e4-b430-4a87-97f5-7b7e5367a692',1,NULL);
INSERT INTO `usergroup` VALUES ('55cda122-44f6-4ce9-a6ea-12ecd4c7701c','2020-05-28 23:55:24.351240657+08:00','2020-05-28 23:55:24.351240657+08:00',NULL,'a22824b3-c6eb-4757-a93c-6bff97337824','f4af0d66-6eff-4420-b95e-b67c6c662739',1,NULL);
INSERT INTO `usergroup` VALUES ('035d0a82-5f8f-47e3-aebd-4d6185eb2ecd','2020-05-28 23:55:24.356319217+08:00','2020-05-28 23:55:24.356319217+08:00',NULL,'28ae1605-73b8-4f9c-b393-d38d2e1addc7','f4af0d66-6eff-4420-b95e-b67c6c662739',1,NULL);
CREATE TABLE IF NOT EXISTS `userfriend` (
	`friend_id`	blob NOT NULL,
	`user_id`	blob NOT NULL,
	PRIMARY KEY(`friend_id`,`user_id`)
);
INSERT INTO `userfriend` VALUES ('3cc8591c-d77f-4715-a48c-b2e0a44e04e4','a22824b3-c6eb-4757-a93c-6bff97337824');
INSERT INTO `userfriend` VALUES ('3cc8591c-d77f-4715-a48c-b2e0a44e04e4','28ae1605-73b8-4f9c-b393-d38d2e1addc7');
INSERT INTO `userfriend` VALUES ('a22824b3-c6eb-4757-a93c-6bff97337824','28ae1605-73b8-4f9c-b393-d38d2e1addc7');
CREATE TABLE IF NOT EXISTS `user` (
	`id`	blob NOT NULL,
	`created_at`	datetime NOT NULL,
	`update_at`	datetime NOT NULL,
	`deleted_at`	datetime,
	`name`	VARCHAR ( 32 ),
	`nickname`	varchar ( 255 ) NOT NULL,
	`gender`	TINYINT DEFAULT 0,
	`avatar`	varchar ( 255 ),
	`mobile`	VARCHAR ( 16 ),
	`email`	varchar ( 255 ) UNIQUE,
	`homepage`	varchar ( 255 ),
	`birthday`	DATE,
	`height`	FLOAT,
	`blood_type`	VARCHAR ( 8 ),
	`notice`	TEXT,
	`intro`	TEXT,
	`address`	JSON,
	`lives`	JSON,
	`tags`	JSON,
	`lucky_numbers`	JSON,
	`score`	integer DEFAULT 0,
	`user_no`	integer PRIMARY KEY AUTOINCREMENT
);
INSERT INTO `user` VALUES ('3cc8591c-d77f-4715-a48c-b2e0a44e04e4','2020-05-28 23:55:24.255227932+08:00','2020-05-28 23:55:24.334525361+08:00',NULL,NULL,'gabe',1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,0,1);
INSERT INTO `user` VALUES ('a22824b3-c6eb-4757-a93c-6bff97337824','2020-05-28 23:55:24.264811919+08:00','2020-05-28 23:55:24.348553082+08:00',NULL,NULL,'jack',0,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,0,2);
INSERT INTO `user` VALUES ('28ae1605-73b8-4f9c-b393-d38d2e1addc7','2020-05-28 23:55:24.269855811+08:00','2020-05-28 23:55:24.354942975+08:00',NULL,NULL,'rose',0,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,0,3);
CREATE TABLE IF NOT EXISTS `rolemenu` (
	`menu_id`	blob NOT NULL,
	`role_id`	blob NOT NULL,
	PRIMARY KEY(`menu_id`,`role_id`)
);
INSERT INTO `rolemenu` VALUES ('5b72d863-ac62-4139-a8fe-e2988f8d37e3','89d17436-f83b-4644-9ebe-feb5ad51e4d9');
INSERT INTO `rolemenu` VALUES ('f784dfb1-af44-4333-bd5d-4ce458726382','89d17436-f83b-4644-9ebe-feb5ad51e4d9');
INSERT INTO `rolemenu` VALUES ('4840529e-dacd-4526-a8c3-7ad47762c2b5','89d17436-f83b-4644-9ebe-feb5ad51e4d9');
INSERT INTO `rolemenu` VALUES ('32730677-54d6-47ce-96f4-904a802f91a2','89d17436-f83b-4644-9ebe-feb5ad51e4d9');
INSERT INTO `rolemenu` VALUES ('5b72d863-ac62-4139-a8fe-e2988f8d37e3','87934398-0bc9-4ffb-98be-18adeefe4967');
INSERT INTO `rolemenu` VALUES ('f784dfb1-af44-4333-bd5d-4ce458726382','87934398-0bc9-4ffb-98be-18adeefe4967');
INSERT INTO `rolemenu` VALUES ('4840529e-dacd-4526-a8c3-7ad47762c2b5','87934398-0bc9-4ffb-98be-18adeefe4967');
CREATE TABLE IF NOT EXISTS `role` (
	`id`	blob NOT NULL,
	`created_at`	datetime NOT NULL,
	`update_at`	datetime NOT NULL,
	`deleted_at`	datetime,
	`name`	VARCHAR ( 32 ) NOT NULL UNIQUE,
	PRIMARY KEY(`id`)
);
INSERT INTO `role` VALUES ('89d17436-f83b-4644-9ebe-feb5ad51e4d9','2020-05-28 23:55:24.273436012+08:00','2020-05-28 23:55:24.33507979+08:00',NULL,'admin');
INSERT INTO `role` VALUES ('87934398-0bc9-4ffb-98be-18adeefe4967','2020-05-28 23:55:24.27535163+08:00','2020-05-28 23:55:24.348963505+08:00',NULL,'user');
INSERT INTO `role` VALUES ('87af54fd-18a8-488c-9438-52d7f032be31','2020-05-28 23:55:24.278521899+08:00','2020-05-28 23:55:24.355532772+08:00',NULL,'guest');
CREATE TABLE IF NOT EXISTS `menu` (
	`id`	blob NOT NULL,
	`created_at`	datetime NOT NULL,
	`update_at`	datetime NOT NULL,
	`deleted_at`	datetime,
	`parent_id`	blob,
	`name`	varchar ( 255 ) NOT NULL,
	`path`	varchar ( 255 ) NOT NULL,
	`icon`	varchar ( 255 ),
	`sort`	integer,
	PRIMARY KEY(`id`)
);
INSERT INTO `menu` VALUES ('5b72d863-ac62-4139-a8fe-e2988f8d37e3','2020-05-28 23:55:24.286790059+08:00','2020-05-28 23:55:24.349222737+08:00',NULL,'00000000-0000-0000-0000-000000000000','welcome','/welcome','dashboard',1);
INSERT INTO `menu` VALUES ('f784dfb1-af44-4333-bd5d-4ce458726382','2020-05-28 23:55:24.290146999+08:00','2020-05-28 23:55:24.349709383+08:00',NULL,'00000000-0000-0000-0000-000000000000','dashboard','/dashboard','dashboard',1);
INSERT INTO `menu` VALUES ('4840529e-dacd-4526-a8c3-7ad47762c2b5','2020-05-28 23:55:24.292545732+08:00','2020-05-28 23:55:24.350124411+08:00',NULL,'f784dfb1-af44-4333-bd5d-4ce458726382','welcome','/dashboard/users',NULL,1);
INSERT INTO `menu` VALUES ('32730677-54d6-47ce-96f4-904a802f91a2','2020-05-28 23:55:24.294338641+08:00','2020-05-28 23:55:24.33722109+08:00',NULL,'f784dfb1-af44-4333-bd5d-4ce458726382','auths','/dashboard/auths',NULL,1);
CREATE TABLE IF NOT EXISTS `group` (
	`id`	blob NOT NULL,
	`created_at`	datetime NOT NULL,
	`update_at`	datetime NOT NULL,
	`deleted_at`	datetime,
	`name`	VARCHAR ( 32 ) NOT NULL,
	`leader_id`	blob NOT NULL,
	PRIMARY KEY(`id`)
);
INSERT INTO `group` VALUES ('f4af0d66-6eff-4420-b95e-b67c6c662739','2020-05-28 23:55:24.282002176+08:00','2020-05-28 23:55:24.356108129+08:00',NULL,'titanic','a22824b3-c6eb-4757-a93c-6bff97337824');
INSERT INTO `group` VALUES ('6dd8d5e4-b430-4a87-97f5-7b7e5367a692','2020-05-28 23:55:24.28464912+08:00','2020-05-28 23:55:24.345654539+08:00',NULL,'rayjar','3cc8591c-d77f-4715-a48c-b2e0a44e04e4');
CREATE TABLE IF NOT EXISTS `auth` (
	`id`	blob NOT NULL,
	`created_at`	datetime NOT NULL,
	`update_at`	datetime NOT NULL,
	`deleted_at`	datetime,
	`user_id`	blob NOT NULL,
	`auth_type`	VARCHAR ( 16 ) NOT NULL,
	`auth_name`	VARCHAR ( 128 ) NOT NULL,
	`auth_code`	varchar ( 255 ),
	`verify_time`	DATETIME,
	`expire_time`	DATETIME,
	`is_enabled`	bool DEFAULT true,
	PRIMARY KEY(`id`)
);
INSERT INTO `auth` VALUES ('ac699fda-9e07-44ed-a3b4-45a67a66d143','2020-05-28 23:55:24.260147864+08:00','2020-05-28 23:55:24.260147864+08:00',NULL,'3cc8591c-d77f-4715-a48c-b2e0a44e04e4','account','gabe','42b83f8b2c3f6ae5360c0bac28ab5233','2020-05-28 23:55:24.259991154+08:00',NULL,1);
CREATE INDEX IF NOT EXISTS `idx_usergroup_deleted_at` ON `usergroup` (
	`deleted_at`
);
CREATE INDEX IF NOT EXISTS `idx_user_deleted_at` ON `user` (
	`deleted_at`
);
CREATE INDEX IF NOT EXISTS `idx_role_deleted_at` ON `role` (
	`deleted_at`
);
CREATE INDEX IF NOT EXISTS `idx_menu_deleted_at` ON `menu` (
	`deleted_at`
);
CREATE INDEX IF NOT EXISTS `idx_group_deleted_at` ON `group` (
	`deleted_at`
);
CREATE INDEX IF NOT EXISTS `idx_auth_deleted_at` ON `auth` (
	`deleted_at`
);
COMMIT;
