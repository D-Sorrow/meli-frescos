-- MySQL dump 10.13  Distrib 8.0.41, for macos15 (arm64)
--
-- Host: 127.0.0.1    Database: bgow15s436
-- ------------------------------------------------------
-- Server version	9.2.0

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
-- Table structure for table `provinces`
--

DROP TABLE IF EXISTS `provinces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `provinces` (
  `id` int NOT NULL AUTO_INCREMENT,
  `province_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `id_country_fk` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_country_fk` (`id_country_fk`),
  CONSTRAINT `provinces_ibfk_1` FOREIGN KEY (`id_country_fk`) REFERENCES `countries` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `provinces`
--

INSERT INTO `provinces` VALUES (1,'Amazonas',1),(2,'Antioquia',1),(3,'Arauca',1),(4,'Atlántico',1),(5,'Bogotá',1),(6,'Bolívar',1),(7,'Boyacá',1),(8,'Caldas',1),(9,'Caquetá',1),(10,'Casanare',1),(11,'Cauca',1),(12,'Cesar',1),(13,'Chocó',1),(14,'Córdoba',1),(15,'Cundinamarca',1),(16,'Guainía',1),(17,'Guaviare',1),(18,'Huila',1),(19,'La Guajira',1),(20,'Magdalena',1),(21,'Meta',1),(22,'Nariño',1),(23,'Norte de Santander',1),(24,'Putumayo',1),(25,'Quindío',1),(26,'Risaralda',1),(27,'San Andrés y Providencia',1),(28,'Santander',1),(29,'Sucre',1),(30,'Tolima',1),(31,'Valle del Cauca',1),(32,'Vaupés',1),(33,'Vichada',1);
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-27 10:34:17
