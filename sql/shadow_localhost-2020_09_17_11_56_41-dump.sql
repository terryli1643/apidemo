-- MySQL dump 10.13  Distrib 5.7.16, for osx10.11 (x86_64)
--
-- Host: 127.0.0.1    Database: shadow
-- ------------------------------------------------------
-- Server version	5.7.16-log

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
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES ('g','ROLE_ANONYMOUS','ROLE_BCC_USER','/bcc',NULL,NULL,NULL),('p','ROLE_BCC_USER','/bcc','/bcc/login','POST','allow','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `id_space`
--

DROP TABLE IF EXISTS `id_space`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `id_space` (
  `space_name` varchar(45) NOT NULL DEFAULT '',
  `prefix` varchar(45) NOT NULL DEFAULT '',
  `suffix` varchar(45) NOT NULL DEFAULT '',
  `seed` bigint(20) unsigned NOT NULL,
  `batch_size` bigint(11) unsigned NOT NULL,
  PRIMARY KEY (`space_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `id_space`
--

LOCK TABLES `id_space` WRITE;
/*!40000 ALTER TABLE `id_space` DISABLE KEYS */;
INSERT INTO `id_space` VALUES ('order','order','WD0ZV6AFGSCU9MBKN24HJL578XIOP31QERTY',78436785006604,1000),('recommendcode','','0ZV6AWDFGSCU9HJL578X1MBKN24QERTYIOP3',78436785643604,1000),('req','','4QEAWDFGSCU9HJRTYI780ZV6L52X1MBKNOP3',78436785643604,1000),('role','ROLE_','',37000,1000),('user','user','',8387000,1000);
/*!40000 ALTER TABLE `id_space` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_admin`
--

DROP TABLE IF EXISTS `t_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `user_type` bigint(20) DEFAULT NULL,
  `account` varchar(256) DEFAULT NULL,
  `job_number` varchar(256) DEFAULT NULL,
  `login_password` varchar(256) DEFAULT NULL,
  `secure_password` varchar(256) DEFAULT NULL,
  `department` varchar(256) DEFAULT NULL,
  `creator` bigint(20) DEFAULT NULL,
  `creator_account` varchar(256) DEFAULT NULL,
  `remark` varchar(256) DEFAULT NULL,
  `state` bigint(20) DEFAULT NULL,
  `last_login_at` datetime DEFAULT NULL,
  `last_login_ip` varchar(256) DEFAULT NULL,
  `last_login_ip_addr` varchar(256) DEFAULT NULL,
  `im_account` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_admin`
--

LOCK TABLES `t_admin` WRITE;
/*!40000 ALTER TABLE `t_admin` DISABLE KEYS */;
INSERT INTO `t_admin` VALUES (1,'2020-08-16 11:41:41','2020-08-16 11:41:41',4,'test','001','$2a$10$BbxBCcMJdXeF5KMNcOdAmONfhVCIcLzfDMILVmWE3ASNGn3DEsPCu','$2a$10$klD7upJamB71pC4p0pm0N.EL1FIC7vzoFLdKj8AZmdqWXqBjmGupq','test department',0,'','',1,NULL,'','','');
/*!40000 ALTER TABLE `t_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_id_space`
--

DROP TABLE IF EXISTS `t_id_space`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_id_space` (
  `space_name` varchar(256) NOT NULL,
  `prefix` varchar(256) DEFAULT NULL,
  `suffix` varchar(256) DEFAULT NULL,
  `seed` bigint(20) DEFAULT NULL,
  `batch_size` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`space_name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_id_space`
--

LOCK TABLES `t_id_space` WRITE;
/*!40000 ALTER TABLE `t_id_space` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_id_space` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role`
--

DROP TABLE IF EXISTS `t_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `role_code` varchar(256) DEFAULT NULL,
  `role_name` varchar(256) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL,
  `state` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role`
--

LOCK TABLES `t_role` WRITE;
/*!40000 ALTER TABLE `t_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role_user_rel`
--

DROP TABLE IF EXISTS `t_role_user_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_role_user_rel` (
  `admin_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`admin_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role_user_rel`
--

LOCK TABLES `t_role_user_rel` WRITE;
/*!40000 ALTER TABLE `t_role_user_rel` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_role_user_rel` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-09-17 11:56:41
