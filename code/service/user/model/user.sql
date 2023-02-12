CREATE TABLE `user_basic`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`         varchar(255) NOT NULL DEFAULT '' ,
    `password`     varchar(255) NOT NULL DEFAULT '',
    `email`        varchar(100) NOT NULL DEFAULT '',
    `now_volume`   int(11) NOT NULL DEFAULT '0' COMMENT '当前存储容量',
    `total_volume` int(11) NOT NULL DEFAULT '1000000000' COMMENT '最大存储容量',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name_unique` (`name`),
    UNIQUE KEY `idx_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
