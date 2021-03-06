-- MySQL Script generated by MySQL Workbench
-- Wed Oct 17 20:51:39 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema wegou
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema wegou
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `wegou` DEFAULT CHARACTER SET utf8 ;
USE `wegou` ;

-- -----------------------------------------------------
-- Table `wegou`.`account`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `wegou`.`account` ;

CREATE TABLE IF NOT EXISTS `wegou`.`account` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '公众号名称',
  `password` CHAR(32) NOT NULL DEFAULT '',
  `created_time` INT UNSIGNED NOT NULL DEFAULT 0,
  `updated_time` INT UNSIGNED NOT NULL DEFAULT 0,
  `login_time` INT NOT NULL DEFAULT 0 COMMENT '上一次登录时间',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态\n1启用\n2禁用',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '用户账号';


-- -----------------------------------------------------
-- Table `wegou`.`wechat`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `wegou`.`wechat` ;

CREATE TABLE IF NOT EXISTS `wegou`.`wechat` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL DEFAULT '',
  `code` VARCHAR(45) NOT NULL DEFAULT '',
  `appid` VARCHAR(255) NOT NULL DEFAULT '',
  `appsecret` VARCHAR(255) NOT NULL DEFAULT '',
  `aeskey` VARCHAR(255) NOT NULL DEFAULT '',
  `oriid` VARCHAR(255) NOT NULL DEFAULT '',
  `token` VARCHAR(255) NOT NULL DEFAULT '',
  `created_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `account_type` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型\\n1 未认证订阅号\\n2 微信认证订阅号\\n3 未认证服务号\\n4 微信认证服务号\\n5 测试公众号',
  `service_type` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '服务类型\\n1 免费服务\\n2 收费服务',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '服务状态\\n1正常\\n2停止',
  `db_host` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '数据库地址',
  `db_name` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '数据库名称',
  `db_port` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '数据库端口',
  `db_user` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '数据库用户名',
  `db_password` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '数据库密码',
  `auth_status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
  `account_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_wechat_account_idx` (`account_id` ASC) )
ENGINE = MyISAM
COMMENT = '公众号';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
