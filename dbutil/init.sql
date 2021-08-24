GRANT ALL PRIVILEGES ON *.* TO 'root' @'%';
GRANT ALL PRIVILEGES ON *.* TO 'airdb' @'%';
FLUSH PRIVILEGES;
CREATE DATABASE `test`;
USE `test`;
CREATE TABLE `user_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;