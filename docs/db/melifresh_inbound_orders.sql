CREATE DATABASE  IF NOT EXISTS `melifresh` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `melifresh`;
-- MySQL dump 10.13  Distrib 8.0.41, for macos15 (arm64)
--
-- Host: altergeistsrveastus2.mysql.database.azure.com    Database: melifresh
-- ------------------------------------------------------
-- Server version	8.0.39-azure

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `inbound_orders`
--

DROP TABLE IF EXISTS `inbound_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inbound_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_date` datetime(6) NOT NULL,
  `order_number` varchar(30) COLLATE utf8mb4_general_ci NOT NULL,
  `employe_id` int NOT NULL,
  `product_batch_id` int NOT NULL,
  `wareHouse_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_number` (`order_number`),
  KEY `employe_id` (`employe_id`),
  KEY `product_batch_id` (`product_batch_id`),
  KEY `wareHouse_id` (`wareHouse_id`),
  CONSTRAINT `inbound_orders_ibfk_1` FOREIGN KEY (`employe_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `inbound_orders_ibfk_2` FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches` (`id`),
  CONSTRAINT `inbound_orders_ibfk_3` FOREIGN KEY (`wareHouse_id`) REFERENCES `warehouses` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inbound_orders`
--

LOCK TABLES `inbound_orders` WRITE;
/*!40000 ALTER TABLE `inbound_orders` DISABLE KEYS */;
INSERT INTO `inbound_orders` VALUES (1,'2023-10-01 10:00:00.000000','ORD0001',1,1,1),(2,'2023-10-02 11:00:00.000000','ORD0002',2,2,2),(3,'2023-10-03 12:00:00.000000','ORD0003',3,3,3),(4,'2023-10-04 13:00:00.000000','ORD0004',4,4,4),(5,'2023-10-05 14:00:00.000000','ORD0005',5,5,5),(6,'2023-10-06 15:00:00.000000','ORD0006',1,5,20),(7,'2023-10-07 16:00:00.000000','ORD0007',2,3,17),(8,'2023-10-08 17:00:00.000000','ORD0008',3,1,12),(9,'2023-10-09 18:00:00.000000','ORD0009',4,2,22),(10,'2023-10-10 19:00:00.000000','ORD0010',5,3,15);
/*!40000 ALTER TABLE `inbound_orders` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:09:51
