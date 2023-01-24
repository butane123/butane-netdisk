CREATE TABLE `user_repository`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `identity`            varchar(36)  DEFAULT NULL,
    `user_identity`       varchar(36)  DEFAULT NULL,
    `parent_id`           int(11)      DEFAULT NULL,
    `repository_identity` varchar(36)  DEFAULT NULL COMMENT '空串则为文件夹',
    `name`                varchar(255) DEFAULT NULL,
    `created_at`          datetime     DEFAULT NULL,
    `updated_at`          datetime     DEFAULT NULL,
    `deleted_at`          datetime     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 8
  DEFAULT CHARSET = utf8;
