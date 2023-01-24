CREATE TABLE `repository_pool`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `identity`   varchar(36)  DEFAULT NULL,
    `hash`       varchar(32)  DEFAULT NULL COMMENT '文件的唯一标识',
    `ext`        varchar(30)  DEFAULT NULL COMMENT '文件扩展名',
    `size`       int(11) DEFAULT NULL COMMENT '文件大小',
    `path`       varchar(255) DEFAULT NULL COMMENT '文件路径',
    `created_at` datetime     DEFAULT NULL,
    `updated_at` datetime     DEFAULT NULL,
    `deleted_at` datetime     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
