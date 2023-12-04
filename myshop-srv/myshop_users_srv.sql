/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : myshop_users_srv

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 05/12/2023 02:06:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `modify_time` datetime(3) DEFAULT NULL,
  `delete_at` datetime(3) DEFAULT NULL,
  `mobile` varchar(11) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nick_name` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `gender` varchar(6) COLLATE utf8mb4_general_ci DEFAULT 'male' COMMENT 'female表示女, male表示男',
  `role` varchar(6) COLLATE utf8mb4_general_ci DEFAULT '1' COMMENT '1表示普通用户, 2表示管理员',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (1, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787860', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman0', NULL, 'male', '2');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (2, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787861', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman1', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (3, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787862', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman2', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (4, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787863', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman3', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (5, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787864', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman4', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (6, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787865', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman5', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (7, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787866', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman6', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (8, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787867', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman7', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (9, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787868', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman8', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (10, '0000-00-00 00:00:00.000', '0000-00-00 00:00:00.000', NULL, '15109787869', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', 'luman9', NULL, 'male', '1');
INSERT INTO `user` (`id`, `add_time`, `modify_time`, `delete_at`, `mobile`, `password`, `nick_name`, `birthday`, `gender`, `role`) VALUES (11, NULL, NULL, NULL, '15509897887', '$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635', NULL, NULL, 'male', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
