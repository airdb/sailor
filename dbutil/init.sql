GRANT ALL PRIVILEGES ON *.* TO 'root' @'%';

CREATE USER 'airdb'@'%' IDENTIFIED BY 'airdb';

ALTER user 'airdb'@'%'IDENTIFIED BY 'airdb';

GRANT ALL PRIVILEGES ON *.* TO 'airdb' @'%';

GRANT ALL ON *.* TO 'airdb'@'%';

FLUSH PRIVILEGES;

CREATE DATABASE IF NOT EXISTS `test`;

USE `test`;

CREATE TABLE
    IF NOT EXISTS `tab_user` (
        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP NULL DEFAULT NULL,
        `user` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;