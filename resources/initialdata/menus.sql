INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (1,'系統管理功能','system:feature',NULL,'T',999,0,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'系統管理功能');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (2,'商品管理功能','production:feature',NULL,'T',998,0,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'商品管理功能');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (3,'使用者管理','user:manage','/backstage/user','P',110,1,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (4,'使用者介面','user','/backstage/user','F',109,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者介面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (5,'使用者新增','user:create','/backstage/user/create','F',109,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (6,'使用者刪除','user:delete','/backstage/user/delete','F',108,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (7,'使用者修改','user:edit','/backstage/user/edit','F',107,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (8,'使用者密碼修改','user:password:edit','/backstage/user/password/edit','F',107,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者密碼修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (9,'使用者密碼重置','user:password:reset','/backstage/user/password/reset','F',107,3,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'使用者密碼重置');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (10,'菜單管理','menu:manage','/backstage/menu','P',120,1,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (11,'菜單介面','menu','/backstage/menu','F',120,10,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單介面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (12,'菜單新增','menu:create','/backstage/menu/create','F',119,10,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (13,'菜單刪除','menu:delete','/backstage/menu/delete','F',118,10,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (14,'菜單修改','menu:edit','/backstage/menu/edit','F',117,10,1,'2021-12-23 17:08:52','2021-12-23 17:08:52',NULL,NULL,NULL,'菜單修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (15,'角色管理','role:manage','/backstage/role','P',150,1,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',1,NULL,NULL,'角色頁面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (16,'角色介面','role','/backstage/role','F',149,15,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色介面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (17,'角色新增','role:create','/backstage/role/create','F',149,15,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (18,'角色刪除','role:delete','/backstage/role/delete','F',148,15,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (19,'角色修改','role:edit','/backstage/role/edit','F',147,15,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'角色修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (20,'商品管理','production:manage','/backstage/production','P',130,2,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (21,'商品介面','production','/backstage/production','F',129,20,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品介面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (22,'商品新增','production:create','/backstage/production/create','F',129,20,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (23,'商品刪除','production:delete','/backstage/production/delete','F',128,20,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (24,'商品修改','production:edit','/backstage/production/edit','F',127,20,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'商品修改');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (25,'cache管理','cache:manage','/backstage/cache','P',138,1,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'cache管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (26,'cache介面','cache','/backstage/cache','F',139,25,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'cache介面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (27,'cache刪除','cache:delete','/backstage/cache/delete','F',138,25,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'cache刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (28,'cache刪除any','cache:delete:any','/backstage/cache/any/delete','F',138,25,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'cache刪除any');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (29,'輪播圖管理','carousel:manage','/backstage/carousel','P',140,2,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'輪播圖管理');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (30,'輪播片介面','carousel','/backstage/carousel','F',139,29,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'輪播圖介面');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (31,'輪播片新增','carousel:create','/backstage/carousel/create','F',139,29,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'輪播圖新增');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (32,'輪播片刪除','carousel:delete','/backstage/carousel/delete','F',138,29,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'輪播圖刪除');
INSERT INTO `menus` (`id`,`name`,`key`,`url`,`feature`,`weight`,`parent`,`status`,`create_time`,`update_time`,`create_user_id`,`update_user_id`,`deleted`,`remark`) VALUES (33,'輪播片修改','carousel:edit','/backstage/carousel/edit','F',138,29,1,'2021-12-23 17:23:19','2021-12-23 17:23:19',NULL,NULL,NULL,'輪播圖修改');