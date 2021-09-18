-- MySQL dump 10.13  Distrib 8.0.23, for osx10.16 (x86_64)
--
-- Host: localhost    Database: go_admin
-- ------------------------------------------------------
-- Server version	5.5.5-10.5.8-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goadmin_menu`
--

DROP TABLE IF EXISTS `goadmin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT 0,
  `type` tinyint(4) unsigned NOT NULL DEFAULT 0,
  `order` int(11) unsigned NOT NULL DEFAULT 0,
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `plugin_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `uuid` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_menu`
--

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;
INSERT INTO `goadmin_menu` VALUES (1,0,1,2,'权限管理','fa-tasks','','','',NULL,'2019-09-09 16:00:00','2021-09-15 12:06:11'),(2,1,1,2,'Users','fa-users','/info/manager',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(3,1,1,3,'Roles','fa-user','/info/roles',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(4,1,1,4,'Permission','fa-ban','/info/permission',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(5,1,1,5,'Menu','fa-bars','/menu',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(6,1,1,6,'Operation log','fa-history','/info/op',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(7,0,1,1,'Dashboard','fa-bar-chart','/application/info','','',NULL,'2019-09-09 16:00:00','2021-09-13 11:37:03'),(8,0,0,9,'文章管理','fa-tasks','','','',NULL,'2021-09-12 13:53:57','2021-09-15 12:05:11'),(9,8,0,10,'发布','fa-edit','/info/posts/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc','','',NULL,'2021-09-12 13:55:53','2021-09-12 14:15:18'),(10,8,0,9,'列表','fa-list-ul','/info/posts','','',NULL,'2021-09-12 13:59:56','2021-09-12 14:15:50'),(11,0,0,11,'标签管理','fa-tasks','','','',NULL,'2021-09-15 11:52:23','2021-09-15 12:05:17'),(12,11,0,12,'添加','fa-edit','/info/tags/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc','','',NULL,'2021-09-15 11:55:16','2021-09-15 11:55:39'),(13,11,0,11,'列表','fa-list-ul','/info/tags','','',NULL,'2021-09-15 11:56:47','2021-09-15 12:04:33'),(14,0,0,7,'页面管理','fa-tasks','/info/pages','','',NULL,'2021-09-15 11:59:53','2021-09-15 12:05:05'),(15,14,0,8,'创建','fa-edit','/info/pages/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc','','',NULL,'2021-09-15 12:00:57','2021-09-15 12:03:59'),(16,14,0,7,'列表','fa-list-ul','/info/pages','','',NULL,'2021-09-15 12:02:09','2021-09-15 12:04:38'),(17,0,0,13,'网站导航','fa-tasks','','','',NULL,'2021-09-15 12:23:51','2021-09-15 12:25:36'),(18,17,0,13,'列表','fa-list-ul','/info/menus','','',NULL,'2021-09-15 12:24:40','2021-09-15 12:25:16'),(19,17,0,18,'创建','fa-edit','/info/menus/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc','','',NULL,'2021-09-15 12:26:19','2021-09-15 12:26:19'),(20,0,0,18,'媒体中心','fa-file-o','/fm/def/list','','',NULL,'2021-09-18 11:10:13','2021-09-18 11:36:18');
/*!40000 ALTER TABLE `goadmin_menu` ENABLE KEYS */;
UNLOCK TABLES;
