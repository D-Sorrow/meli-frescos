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
-- Table structure for table `purchase_orders`
--

DROP TABLE IF EXISTS `purchase_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `purchase_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_number` varchar(30) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `order_date` datetime(6) DEFAULT NULL,
  `tracking_code` char(36) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `buyer_id` int NOT NULL,
  `carrier_id` int NOT NULL,
  `order_status_id` int NOT NULL,
  `wareHouse_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tracking_code` (`tracking_code`),
  UNIQUE KEY `order_number` (`order_number`),
  KEY `buyer_id` (`buyer_id`),
  KEY `order_status_id` (`order_status_id`),
  KEY `carrier_id` (`carrier_id`),
  KEY `fk_wareHouse` (`wareHouse_id`),
  CONSTRAINT `fk_wareHouse` FOREIGN KEY (`wareHouse_id`) REFERENCES `warehouses` (`id`),
  CONSTRAINT `purchase_orders_ibfk_1` FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`),
  CONSTRAINT `purchase_orders_ibfk_2` FOREIGN KEY (`order_status_id`) REFERENCES `order_status` (`id`),
  CONSTRAINT `purchase_orders_ibfk_3` FOREIGN KEY (`carrier_id`) REFERENCES `carriers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchase_orders`
--

LOCK TABLES `purchase_orders` WRITE;
/*!40000 ALTER TABLE `purchase_orders` DISABLE KEYS */;
INSERT INTO `purchase_orders` VALUES (3,'OR0003','2025-02-15 04:08:52.000000','d2af4044-3c92-4d0b-9de6-9678c755f6c1',1,1,1,1),(4,'OR0004','2025-02-15 04:09:03.000000','d07a0493-0882-4702-b3aa-2dfe9ae108fe',1,1,1,1);
/*!40000 ALTER TABLE `purchase_orders` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`devfresh`@`%`*/ /*!50003 TRIGGER `before_insert_purchase_orders` BEFORE INSERT ON `purchase_orders` FOR EACH ROW BEGIN
    DECLARE next_id INT;
    
    SELECT AUTO_INCREMENT INTO next_id
    FROM information_schema.TABLES
    WHERE TABLE_NAME = 'purchase_orders'
    AND TABLE_SCHEMA = DATABASE();
    
    SET NEW.order_number = CONCAT('OR', LPAD(next_id, 4, '0'));
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:10:01
