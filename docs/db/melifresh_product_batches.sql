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
-- Table structure for table `product_batches`
--

DROP TABLE IF EXISTS `product_batches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_batches` (
  `id` int NOT NULL AUTO_INCREMENT,
  `batch_number` varchar(30) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `current_quantity` int NOT NULL,
  `current_temperature` decimal(19,2) NOT NULL,
  `due_date` datetime(6) NOT NULL,
  `initial_quantity` int NOT NULL,
  `manufacturing_date` datetime(6) NOT NULL,
  `manufacturing_hour` datetime(6) NOT NULL,
  `minimum_temperature` decimal(19,2) NOT NULL,
  `product_id` int NOT NULL,
  `section_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  KEY `section_id` (`section_id`),
  CONSTRAINT `product_batches_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `product_batches_ibfk_2` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_batches`
--

LOCK TABLES `product_batches` WRITE;
/*!40000 ALTER TABLE `product_batches` DISABLE KEYS */;
INSERT INTO `product_batches` VALUES (1,'BATCH001',75,5.00,'2023-12-01 00:00:00.000000',100,'2023-10-01 00:00:00.000000','2023-10-01 08:00:00.000000',0.50,1,1),(2,'BATCH002',50,4.50,'2023-12-15 00:00:00.000000',80,'2023-09-15 00:00:00.000000','2023-09-15 09:00:00.000000',1.00,5,2),(3,'BATCH003',25,6.00,'2023-11-20 00:00:00.000000',30,'2023-09-01 00:00:00.000000','2023-09-01 10:00:00.000000',1.50,4,1),(4,'BATCH004',100,3.50,'2023-11-30 00:00:00.000000',120,'2023-08-20 00:00:00.000000','2023-08-20 07:30:00.000000',0.00,2,2),(5,'BATCH005',10,2.50,'2023-10-10 00:00:00.000000',10,'2023-07-12 00:00:00.000000','2023-07-12 14:30:00.000000',1.00,1,1),(6,'BATCH006',10,2.50,'2023-10-10 00:00:00.000000',10,'2023-07-12 00:00:00.000000','2023-07-12 14:30:00.000000',1.00,7,1);
/*!40000 ALTER TABLE `product_batches` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:10:09
