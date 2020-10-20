SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for reminder
-- ----------------------------
DROP TABLE IF EXISTS `reminder`;
CREATE TABLE `reminder`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `note` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '事件记录',
  `alert` int(11) NULL DEFAULT NULL COMMENT '提醒时间',
  `yinli` int(1) UNSIGNED ZEROFILL NOT NULL COMMENT '0-yangli,1-yinli',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- View structure for yangli
-- ----------------------------
DROP VIEW IF EXISTS `yangli`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `yangli` AS select date_format((`reminder`.`date` - interval `reminder`.`alert` day),'%m%d') AS `startdate`,date_format(`reminder`.`date`,'%m%d') AS `stopdate`,`reminder`.`note` AS `note` from `reminder` where (`reminder`.`yinli` = 0);

-- ----------------------------
-- View structure for yinli
-- ----------------------------
DROP VIEW IF EXISTS `yinli`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `yinli` AS select date_format((`reminder`.`date` - interval `reminder`.`alert` day),'%m%d') AS `startdate`,date_format(`reminder`.`date`,'%m%d') AS `stopdate`,`reminder`.`note` AS `note` from `reminder` where (`reminder`.`yinli` = 1);

SET FOREIGN_KEY_CHECKS = 1;
