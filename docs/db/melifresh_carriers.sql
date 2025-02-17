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
-- Table structure for table `carriers`
--

DROP TABLE IF EXISTS `carriers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `carriers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cid` varchar(30) COLLATE utf8mb4_general_ci NOT NULL,
  `company_name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `address` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `telephone` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `locality_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `cid` (`cid`),
  KEY `locality_id` (`locality_id`),
  CONSTRAINT `carriers_ibfk_1` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carriers`
--

LOCK TABLES `carriers` WRITE;
/*!40000 ALTER TABLE `carriers` DISABLE KEYS */;
INSERT INTO `carriers` VALUES (1,'CARRIER001','Transporte Rápido S.A.','Av. Libertador 1234, Buenos Aires','+54 11 1234 5678',1),(2,'CARRIER002','Logística Express','Calle 45, Sector Industrial, Montevideo','+598 2 2345 6789',2),(3,'CARRIER003','Servicios de Carga Segura','Av. Principal 987, Lima','+51 1 3456 7890',3),(4,'CARRIER004','Red de Transporte Internacional','Calle Fuerte 456, Quito','+593 2 4567 8901',4),(5,'CARRIER005','Transporte Global','Boulevard 1500, Santiago','+56 2 5678 9012',5),(6,'CARRIER006','Transporte Ágil','Calle 5, Medellín','+57 4 6789 0123',6),(7,'CARRIER007','Logística Andina','Avenida Andes 89, Bogotá','+57 1 7890 1234',7),(8,'CARRIER008','Carga Rápida','Calle 34, Caracas','+58 212 8901 2345',8),(9,'CARRIER009','Transporte del Pacífico','Av. Costera 999, Guayaquil','+593 4 5678 9012',9),(10,'CARRIER010','Logística Moderna','Calle El Sol 123, Buenos Aires','+54 9 11 1234 5678',10),(11,'CARRIER011','Transportes del Sur','Calle Florida 100, Asunción','+595 21 2345 6789',11),(12,'CARRIER012','Servicios Logísticos Fast','Av. Brasil 4321, Montevideo','+598 2 1234 5678',12),(13,'CARRIER013','Cargas Internacionales','Calle Japón 567, Santiago','+56 9 2345 6789',13),(14,'CARRIER014','Transporte Intercontinental','Av. El Lago 3210, Lima','+51 1 8765 4321',14),(15,'CARRIER015','Red Logística Global','Avenida Paz 654, Quito','+593 2 8901 2345',15),(16,'CARRIER016','Logística Eficiente','Calle Central 123, Medellín','+57 4 5678 9012',16),(17,'CARRIER017','Red de Carga Andina','Avenida Libertador 100, Bogotá','+57 1 2345 6789',17),(18,'CARRIER018','Logística Global Express','Calle O''Higgins 789, Caracas','+58 212 3456 7890',18),(19,'CARRIER019','Carga Urgente','Av. de la Costa 456, Guayaquil','+593 4 2345 6789',19),(20,'CARRIER020','Servicios Logísticos Integrados','Calle Montevideo 111, Asunción','+595 21 8765 4321',20),(21,'CARRIER021','Transporte Eficiente','Av. Constitución 876, Buenos Aires','+54 11 2345 6789',21),(22,'CARRIER022','Logística Ágil','Calle 23, Lima','+51 1 9876 5432',22),(23,'CARRIER023','Cargas Express S.A.','Av. Libertad 890, Quito','+593 2 5432 1098',23);
/*!40000 ALTER TABLE `carriers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:09:56
