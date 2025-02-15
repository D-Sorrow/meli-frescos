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
-- Table structure for table `product_records`
--

DROP TABLE IF EXISTS `product_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_records` (
  `id` int NOT NULL AUTO_INCREMENT,
  `last_update_date` datetime(6) NOT NULL,
  `purchase_price` decimal(19,2) NOT NULL,
  `sale_price` decimal(19,2) NOT NULL,
  `product_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `product_records_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_records`
--

LOCK TABLES `product_records` WRITE;
/*!40000 ALTER TABLE `product_records` DISABLE KEYS */;
INSERT INTO `product_records` VALUES (1,'2025-02-12 10:30:00.000000',12.50,19.99,1),(2,'2025-02-12 10:30:00.000000',6.75,12.50,2),(3,'2025-02-12 10:30:00.000000',15.99,25.00,3),(4,'2025-02-12 10:30:00.000000',8.50,14.99,4),(5,'2025-02-12 10:30:00.000000',3.99,7.50,5),(6,'2025-02-12 10:30:00.000000',9.99,18.50,6),(7,'2025-02-12 10:30:00.000000',4.50,8.99,7),(8,'2025-02-12 10:30:00.000000',10.99,17.50,8),(9,'2025-02-12 10:30:00.000000',22.75,35.00,9),(10,'2025-02-12 10:30:00.000000',19.00,29.99,10),(11,'2025-02-12 10:30:00.000000',15.50,22.00,11),(12,'2025-02-12 10:30:00.000000',4.75,9.99,12),(13,'2025-02-12 10:30:00.000000',12.00,18.99,13),(14,'2025-02-12 10:30:00.000000',3.25,6.99,14),(15,'2025-02-12 10:30:00.000000',5.50,11.99,15),(16,'2025-02-12 10:30:00.000000',9.75,17.99,16),(17,'2025-02-12 10:30:00.000000',18.50,28.99,17),(18,'2025-02-12 10:30:00.000000',14.99,22.50,18),(19,'2025-02-12 10:30:00.000000',5.50,9.99,19),(20,'2025-02-12 10:30:00.000000',10.00,15.99,20),(21,'2025-02-12 10:30:00.000000',8.25,14.99,21),(22,'2025-02-12 10:30:00.000000',6.50,12.00,22),(23,'2025-02-12 10:30:00.000000',12.50,20.99,23),(24,'2025-02-12 10:30:00.000000',4.25,8.50,24),(25,'2025-02-12 10:30:00.000000',7.00,13.99,25),(26,'2025-02-12 10:30:00.000000',9.99,19.99,26),(27,'2025-02-12 10:30:00.000000',15.50,29.99,27),(28,'2025-02-12 10:30:00.000000',6.75,14.50,28),(29,'2025-02-12 10:30:00.000000',3.99,8.50,29),(30,'2025-02-12 10:30:00.000000',7.25,15.99,30),(31,'2025-02-12 10:30:00.000000',4.99,10.99,31),(32,'2025-02-12 10:30:00.000000',12.75,22.99,32),(33,'2025-02-12 10:30:00.000000',9.00,18.50,33),(34,'2025-02-12 10:30:00.000000',14.00,29.99,34),(35,'2025-02-12 10:30:00.000000',18.50,36.50,35),(36,'2025-02-12 10:30:00.000000',10.50,21.99,36),(37,'2025-02-12 10:30:00.000000',11.25,23.99,37),(38,'2025-02-12 10:30:00.000000',5.99,12.99,38),(39,'2025-02-12 10:30:00.000000',3.99,8.50,39),(40,'2025-02-12 10:30:00.000000',7.00,14.50,40),(41,'2025-02-12 10:30:00.000000',5.50,11.99,41),(42,'2025-02-12 10:30:00.000000',6.25,12.00,42),(43,'2025-02-12 10:30:00.000000',3.50,7.99,43),(44,'2025-02-12 10:30:00.000000',9.00,19.99,44),(45,'2025-02-12 10:30:00.000000',12.50,22.50,45),(46,'2025-02-12 10:30:00.000000',7.75,16.50,46),(47,'2025-02-12 10:30:00.000000',8.25,17.00,47),(48,'2025-02-12 10:30:00.000000',6.00,13.50,48),(49,'2025-02-12 10:30:00.000000',4.25,8.99,49),(50,'2025-02-12 10:30:00.000000',5.00,10.99,50),(51,'2025-02-12 10:30:00.000000',3.99,7.99,51),(52,'2025-02-12 10:30:00.000000',2.50,5.50,52),(53,'2025-02-12 10:30:00.000000',4.20,8.00,53),(54,'2025-02-12 10:30:00.000000',6.50,12.00,54),(55,'2025-02-12 10:30:00.000000',5.00,9.50,55),(56,'2025-02-12 10:30:00.000000',6.50,12.50,56),(57,'2025-02-12 10:30:00.000000',8.00,16.99,57),(58,'2025-02-12 10:30:00.000000',4.50,9.99,58),(59,'2025-02-12 10:30:00.000000',5.00,10.50,59),(60,'2025-02-12 10:30:00.000000',7.50,14.00,60),(61,'2025-02-12 10:30:00.000000',10.00,20.99,61),(62,'2025-02-12 10:30:00.000000',15.00,29.99,62),(63,'2025-02-12 10:30:00.000000',20.00,39.99,63),(64,'2025-02-12 10:30:00.000000',7.50,14.50,64),(65,'2025-02-12 10:30:00.000000',10.00,19.99,65),(66,'2025-02-12 10:30:00.000000',5.99,12.99,66),(67,'2025-02-12 10:30:00.000000',25.99,49.99,67),(68,'2025-02-12 10:30:00.000000',12.50,19.99,68),(69,'2025-02-12 10:30:00.000000',18.00,34.99,69),(70,'2025-02-12 10:30:00.000000',9.00,18.50,70);
/*!40000 ALTER TABLE `product_records` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:10:12
