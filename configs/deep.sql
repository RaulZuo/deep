CREATE DATABASE IF NOT EXISTS `deep` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `deep`;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `insight_data`
--

DROP TABLE IF EXISTS `insight_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `insight_data` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `gmt_create` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `gmt_modified` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '修改时间',
  `insight_id` varchar(32) NOT NULL COMMENT '洞察数据实例ID uid',
  `name` varchar(100) NOT NULL COMMENT '洞察数据名称',
  `type` varchar(32) NOT NULL DEFAULT 'CONNECT' COMMENT '洞察数据类型',
  `config` json DEFAULT NULL COMMENT '洞察数据配置信息',
  `status` tinyint(1) unsigned NOT NULL DEFAULT 1 COMMENT '逻辑删除，1: online, 2: delete',
  `creator` bigint(20) unsigned NOT NULL COMMENT '创建人ID',
  `creator_name` varchar(100) DEFAULT NULL COMMENT '创建人',
  `modifier` bigint(20) unsigned NOT NULL COMMENT '修改人ID',
  `modifier_name` varchar(100) DEFAULT NULL COMMENT '修改人',
  `process_status` varchar(32) NOT NULL DEFAULT 'RUNNING' COMMENT '数据处理状态',
  `process_instance_id` varchar(100) DEFAULT NULL COMMENT '异步数据处理ID',
  `project_id` varchar(32) NOT NULL COMMENT '洞察数据所属项目ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_insight_id` (`insight_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8 COMMENT='洞察数据实例表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `insight_meta_field`
--

DROP TABLE IF EXISTS `insight_meta_field`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `insight_meta_field` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `insight_id` varchar(32) NOT NULL COMMENT '数据洞察实例ID uid',
    `meta_type` varchar(32) NOT NULL DEFAULT 'TABLE' COMMENT '元字段类型',
    `survey_id` int(10) unsigned DEFAULT NULL COMMENT '问卷ID',
    `survey_version` int(10) unsigned DEFAULT NULL COMMENT '问卷版本',
    `question_code` varchar(32) DEFAULT NULL COMMENT '问卷题目code，合并问卷时为对应的field_code',
    `sub_question_code` varchar(32) DEFAULT NULL COMMENT '问卷子题目code',
    `table_name` varchar(100) DEFAULT NULL COMMENT '表名',
    `table_key` varchar(100) DEFAULT NULL COMMENT '表字段key',
    `table_key_type` varchar(32) DEFAULT NULL COMMENT '表字段key对应的类型',
    `title` varchar(100) NOT NULL COMMENT '元字段名称',
    `visible` tinyint(1) unsigned NOT NULL DEFAULT 1 COMMENT '字段是否可见',
    `status` tinyint(1) unsigned NOT NULL DEFAULT 1 COMMENT '逻辑删除，1: online, 2: delete',
    `calculate_type` varchar(32) NOT NULL COMMENT '字段计算口径',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_field` (`insight_id`, `survey_id`, `survey_version`, `question_code`, `sub_question_code`, `table_name`, `table_key`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8 COMMENT '元字段表';
/*!40101 SET character_set_client = @saved_cs_client */;
