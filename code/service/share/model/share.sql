CREATE TABLE `share_basic` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT '0',
  `repository_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '公共池中的唯一标识',
  `user_repository_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户池子中的唯一标识',
  `expired_time` int(11) NOT NULL DEFAULT '0' COMMENT '失效时间，单位秒, 【0-永不失效】',
  `click_num` int(11) NOT NULL DEFAULT '0' COMMENT '点击次数',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_repository_id` (`repository_id`),
  KEY `idx_user_repository_id` (`user_repository_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
