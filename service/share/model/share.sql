CREATE TABLE `share_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL COMMENT '公共池中的唯一标识',
  `user_repository_identity` varchar(36) DEFAULT NULL COMMENT '用户池子中的唯一标识',
  `expired_time` int(11) DEFAULT NULL COMMENT '失效时间，单位秒, 【0-永不失效】',
  `click_num` int(11) DEFAULT '0' COMMENT '点击次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;