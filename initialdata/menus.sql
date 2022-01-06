INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (1,'系統管理功能','system:feature',NULL,'T',999,0,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'系統管理功能');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (2,'商品管理功能','production:feature',NULL,'T',998,0,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'商品管理功能');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (3,'使用者管理','user:manage','/backstage/user','P',110,1,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (4,'使用者新增','user:create','/backstage/user/create','F',109,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (5,'使用者刪除','user:delete','/backstage/user/delete','F',108,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (6,'使用者修改','user:update','/backstage/user/update','F',107,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (7,'菜單管理','menu:manage','/backstage/menu','P',120,1,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (8,'菜單新增','menu:create','/backstage/menu/create','F',119,7,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (9,'菜單刪除','menu:delete','/backstage/menu/delete','F',118,7,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (10,'菜單修改','menu:update','/backstage/menu/update','F',117,7,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (11,'商品管理','production:manage','/backstage/production','P',130,2,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (12,'商品新增','production:create','/backstage/production/create','F',129,11,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (13,'商品刪除','production:delete','/backstage/production/delete','F',128,11,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (14,'商品修改','production:update','/backstage/production/update','F',127,11,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (15,'圖片管理','picture:manage','/backstage/picture','P',140,2,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'圖片管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (16,'圖片新增','picture:create','/backstage/picture/create','F',139,15,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'圖片新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (17,'圖片刪除','picture:delete','/backstage/picture/delete','F',138,15,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'圖片刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (18,'角色頁面','role:manage','/backstage/role','P',150,1,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',1,NULL,NULL,'角色頁面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (19,'角色新增','role:create','/backstage/role/create','F',149,18,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (20,'角色刪除','role:delete','/backstage/role/delete','F',148,18,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (21,'角色修改','role:update','/backstage/role/update','F',147,18,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色修改');