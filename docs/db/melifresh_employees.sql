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
-- Table structure for table `employees`
--

DROP TABLE IF EXISTS `employees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employees` (
  `id` int NOT NULL,
  `card_number_id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `last_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `warehouse_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `card_number_id` (`card_number_id`),
  KEY `warehouse_id` (`warehouse_id`),
  CONSTRAINT `employees_ibfk_1` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (1,'1234','Jordan2','Hernandez1',2),(2,'12345678901234567891','Nombre2','Apellido2',2),(3,'12345678901234567892','Nombre3','Apellido3',3),(4,'12345678901234567893','Nombre4','Apellido4',4),(5,'12345678901234567894','Nombre5','Apellido5',5),(6,'12345678901234567895','Nombre6','Apellido6',6),(7,'12345678901234567896','Nombre7','Apellido7',7),(8,'12345678901234567897','Nombre8','Apellido8',8),(9,'12345678901234567898','Nombre9','Apellido9',9),(10,'12345678901234567899','Nombre10','Apellido10',10),(11,'12345678901234567900','Nombre11','Apellido11',11),(12,'12345678901234567901','Nombre12','Apellido12',12),(13,'12345678901234567902','Nombre13','Apellido13',13),(14,'12345678901234567903','Nombre14','Apellido14',14),(15,'12345678901234567904','Nombre15','Apellido15',15),(16,'12345678901234567905','Nombre16','Apellido16',16),(17,'12345678901234567906','Nombre17','Apellido17',17),(18,'12345678901234567907','Nombre18','Apellido18',18),(19,'12345678901234567908','Nombre19','Apellido19',19),(20,'12345678901234567909','Nombre20','Apellido20',20),(21,'12345678901234567910','Nombre21','Apellido21',4),(22,'12345678901234567911','Nombre22','Apellido22',22),(23,'12345678901234567912','Nombre23','Apellido23',22),(24,'12345678901234567913','Nombre24','Apellido24',1),(25,'12345678901234567914','Nombre25','Apellido25',6),(26,'12345678901234567915','Nombre26','Apellido26',27),(27,'12345678901234567916','Nombre27','Apellido27',1),(28,'12345678901234567917','Nombre28','Apellido28',1),(29,'12345678901234567918','Nombre29','Apellido29',2),(30,'12345678901234567919','Nombre30','Apellido30',3);
/*!40000 ALTER TABLE `employees` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:10:05
