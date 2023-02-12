CREATE TABLE `repository_pool`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `hash`       varchar(32) NOT NULL DEFAULT '' COMMENT '文件的唯一标识',
    `ext`        varchar(30) NOT NULL DEFAULT '' COMMENT '文件扩展名',
    `size`       int(11) NOT NULL DEFAULT '0' COMMENT '文件大小',
    `path`       varchar(255) NOT NULL DEFAULT '' COMMENT '文件路径',
    `name`       varchar(255) NOT NULL DEFAULT '',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_hash_unique` (`hash`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
