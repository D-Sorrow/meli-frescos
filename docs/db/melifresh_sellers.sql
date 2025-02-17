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
-- Table structure for table `sellers`
--

DROP TABLE IF EXISTS `sellers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sellers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cid` int NOT NULL,
  `company_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `address` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `telephone` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `locality_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `cid` (`cid`),
  KEY `locality_id` (`locality_id`),
  CONSTRAINT `sellers_ibfk_1` FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sellers`
--

LOCK TABLES `sellers` WRITE;
/*!40000 ALTER TABLE `sellers` DISABLE KEYS */;
INSERT INTO `sellers` VALUES (1,2044,'Hintz Group','Apt 252','+86-872-917-3433',14),(2,5544,'Jakubowski Group','Room 1112','+358-507-307-0538',2),(3,996,'Hilll, McCullough and Quitzon','7th Floor','+62-830-606-6313',3),(4,4315,'Rowe and Sons','18th Floor','+86-712-540-5302',4),(5,5438,'Kunze-Dickinson','PO Box 39664','+86-984-250-0340',5),(6,2637,'Hayes-Block','Apt 211','+420-865-701-6082',6),(7,8997,'Torphy-Brekke','Suite 8','+507-569-781-1995',7),(8,404,'Buckridge, Nolan and Mann','Suite 25','+351-214-503-9042',8),(9,9890,'Lakin Inc','7th Floor','+7-115-526-2469',9),(10,2433,'Carroll, Grimes and Windler','Suite 51','+62-656-917-2374',10),(11,4993,'Runolfsdottir and Sons','Apt 1249','+255-233-548-2230',11),(12,3116,'Kemmer, Reinger and Borer','Room 824','+593-490-175-5458',12),(13,6472,'Murphy Inc','PO Box 10943','+33-575-643-9921',13),(14,3672,'Sipes, Muller and Tremblay','Apt 1177','+86-216-250-0702',14),(15,7382,'Goldner-Bogan','Suite 56','+33-468-455-9278',15),(16,4464,'West Inc','Room 373','+86-942-722-4823',16),(17,7105,'Schowalter and Sons','20th Floor','+880-112-937-3098',17),(18,4885,'Terry LLC','PO Box 53808','+591-362-595-5905',18),(19,9280,'Mraz Inc','PO Box 82391','+86-278-836-2191',19),(20,1654,'Stiedemann and Sons','Suite 85','+86-786-265-5523',20),(21,4272,'Spencer Group','Suite 56','+351-198-769-6841',21),(22,9065,'Windler Group','Apt 1090','+84-997-699-7861',22),(23,4075,'Adams, Blanda and Kessler','Room 1308','+86-992-971-2959',23),(24,9665,'Davis-Bode','2nd Floor','+63-759-792-7539',13),(25,4890,'Altenwerth LLC','Apt 1733','+47-948-540-0704',15),(26,4106,'Kreiger and Sons','Suite 61','+351-200-273-5695',20),(27,997,'VonRueden Group','Room 1310','+62-657-195-3351',9),(28,9005,'Hegmann and Sons','Apt 1167','+45-984-964-9485',6),(29,975,'Nitzsche LLC','PO Box 32770','+60-961-401-2612',4),(30,169,'Robel Group','20th Floor','+850-370-166-7170',8),(31,4643,'Hane, Bergstrom and Stehr','Suite 98','+7-605-815-1083',19),(32,3738,'Dibbert, Murphy and Bins','16th Floor','+967-714-496-7329',11),(33,9024,'Jacobi, Sanford and Larson','Suite 76','+222-663-153-6353',17),(34,6712,'Boehm, Graham and Douglas','Room 638','+358-579-571-0683',19),(35,7404,'Tromp-Tremblay','6th Floor','+86-134-786-6425',19),(36,1567,'Gerlach and Sons','Apt 769','+691-750-858-0875',4),(37,7519,'Lehner LLC','Room 732','+46-430-850-2834',14),(38,1694,'Willms-Johnston','Suite 95','+86-145-130-7894',14),(39,1956,'Kub-Wisoky','Room 128','+55-416-574-2531',17),(40,9393,'Anderson-Heidenreich','12th Floor','+62-104-431-6109',9),(41,1545,'Feil Inc','Apt 641','+86-553-549-0876',22),(42,2389,'Rodriguez-Roob','Apt 1965','+351-464-861-4884',4),(43,3710,'Dicki, Satterfield and Kertzmann','Room 139','+81-546-441-5632',22),(44,7493,'Berge-Wuckert','PO Box 19750','+595-816-932-5500',21),(45,3554,'Kerluke-Cassin','Room 162','+86-773-375-3580',21),(46,9326,'Tremblay Group','Room 1810','+62-672-966-2289',8),(47,5944,'Swaniawski, Hintz and McKenzie','1st Floor','+62-275-398-6847',22),(48,649,'Feeney, Reichert and Hackett','Room 1264','+86-466-183-9851',4),(49,226,'Kunde-Reynolds','PO Box 85417','+81-813-226-2259',11),(50,9381,'Kunze-Lakin','15th Floor','+64-193-626-5483',18),(51,7644,'Frami-Hoppe','18th Floor','+92-112-429-2170',18),(52,2551,'Ortiz-Volkman','Suite 79','+53-895-906-6104',10),(53,445,'Mante-Buckridge','Suite 50','+7-402-701-2505',5),(54,3412,'Stokes LLC','PO Box 3011','+963-375-723-9329',11),(55,283,'Hartmann, Murray and Hirthe','PO Box 49871','+381-114-876-3477',3),(56,7555,'Blick-Bashirian','Room 427','+30-547-639-8057',4),(57,6104,'Jerde-Sanford','Suite 33','+86-636-720-7136',2),(58,4728,'Langworth-Bashirian','Room 1965','+420-881-488-0554',15),(59,8207,'Cronin, Prosacco and Roob','PO Box 83354','+81-886-253-0313',10),(60,244,'Schoen-Deckow','Suite 20','+234-364-273-6332',4),(61,1004,'Emard-Tillman','PO Box 63794','+62-969-638-4995',23),(62,1190,'Macejkovic-Hintz','20th Floor','+86-360-430-7667',18),(63,4047,'Howe, Mayert and Schowalter','Room 1613','+86-713-164-3267',6),(64,1142,'Daniel, Jenkins and Buckridge','Suite 25','+86-900-488-2835',3),(65,2221,'Labadie Inc','Apt 509','+1-345-365-5836',11),(66,4975,'Wunsch-Crooks','9th Floor','+1-465-748-9978',17),(67,6829,'Trantow-Corkery','Suite 74','+86-299-423-4125',20),(68,8273,'Botsford-Bayer','16th Floor','+374-735-898-9252',9),(69,9646,'Cummerata-Okuneva','19th Floor','+995-785-829-9348',13),(70,4517,'Bernhard, Leuschke and Treutel','PO Box 98872','+86-749-923-2097',6),(71,7074,'Yundt-Kuhic','Suite 14','+52-767-182-1544',19),(72,5757,'Huels, Denesik and Lehner','PO Box 85391','+966-944-627-4409',21),(73,5337,'Russel LLC','Suite 48','+55-473-142-9286',21),(74,877,'O''Reilly and Sons','PO Box 60656','+86-840-749-2402',20),(75,5301,'O''Kon Inc','PO Box 56659','+66-900-439-8687',15),(76,3822,'Murazik-Tromp','8th Floor','+62-185-307-0780',3),(77,3453,'Tillman-Abshire','Room 1583','+54-391-966-8579',15),(78,3329,'Williamson Group','Room 397','+62-370-501-7221',14),(79,6381,'Reichert-Kozey','Room 199','+375-778-397-8639',14),(80,309,'Haag-Heathcote','Apt 953','+86-417-104-3729',7),(81,5465,'DuBuque, Hilpert and Hayes','18th Floor','+51-500-236-0218',22),(82,4381,'Jacobson, Gibson and Kohler','Suite 63','+62-609-684-7013',19),(83,8974,'Hoeger-Homenick','PO Box 12923','+33-600-846-2998',6),(84,4173,'Rodriguez and Sons','Suite 5','+86-878-778-6235',10),(85,5816,'DuBuque and Sons','PO Box 10971','+63-592-217-2120',1),(86,172,'Wunsch-Anderson','7th Floor','+358-842-826-2661',8),(87,6654,'Ziemann, Kohler and Mosciski','16th Floor','+222-286-296-2064',13),(88,1115,'Shields, Rowe and Bashirian','PO Box 21995','+7-300-152-0604',12),(89,7008,'Shanahan LLC','Suite 62','+55-242-546-0884',7),(90,8002,'Pfeffer-Maggio','Suite 58','+1-213-261-9628',23),(91,8484,'Smitham, Sawayn and Borer','Apt 244','+62-934-156-0642',6),(92,1403,'Heaney-Block','Suite 71','+30-370-130-0600',8),(93,8211,'Jast Group','Room 854','+86-987-384-2162',17),(94,1493,'Robel, Wunsch and Bashirian','Room 13','+33-546-465-0730',19),(95,2052,'Pfannerstill Inc','Suite 55','+51-387-926-0979',4),(96,7549,'Stracke-Larkin','Suite 42','+52-886-620-2518',17),(97,4118,'Hills LLC','Suite 1','+86-364-919-3962',22),(98,1366,'Senger, Runte and Grady','Suite 33','+86-297-837-6988',19),(99,5343,'Deckow Group','Room 1033','+47-697-633-1680',10),(100,8994,'Mayert and Sons','Suite 7','+7-380-811-1562',7);
/*!40000 ALTER TABLE `sellers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:09:43
