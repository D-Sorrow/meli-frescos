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
-- Table structure for table `buyers`
--

DROP TABLE IF EXISTS `buyers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buyers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `card_number_id` varchar(12) COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `last_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `card_number_id` (`card_number_id`)
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buyers`
--

LOCK TABLES `buyers` WRITE;
/*!40000 ALTER TABLE `buyers` DISABLE KEYS */;
INSERT INTO `buyers` VALUES (1,'D1288325046','Leonore','Brennans'),(2,'S1900529063','Loella','O''Skehan'),(3,'O1051882578','Erina','Sirl'),(4,'C1045449348','Jesus','Seadon'),(5,'S1589446379','Beret','Foch'),(6,'O1157879189','Nobie','Olligan'),(7,'A1166605019','Hermina','Maffini'),(8,'T1695585521','Dyan','Bowyer'),(9,'T1404140128','Colet','Eslie'),(10,'B1858996865','Thaddus','Absalom'),(11,'Z1503543663','Kathryne','Mergue'),(12,'R1495551461','Antony','Shanks'),(13,'P1966539851','Cary','Carlton'),(14,'Y1900286783','Stephine','Lydster'),(15,'T1404164619','Jedd','Grafham'),(16,'Y1188236299','Lissa','Jiracek'),(17,'T1882136203','Sheffield','Carayol'),(18,'J1719451299','Wynne','Dowderswell'),(19,'R1447007729','Franciskus','Barwick'),(20,'Q1990956285','Brooks','Duddell'),(21,'V1666435393','Brooks','Gorrick'),(22,'E1920371559','Sher','Pimley'),(23,'Z1441462180','Abbie','Gallaher'),(24,'Q1585309984','Phillida','Briscoe'),(25,'R1745869283','Clem','Prosser'),(26,'D1424732106','Virge','Dollard'),(27,'U1847451401','Annetta','Derell'),(28,'N1450025292','Kaylyn','Castells'),(29,'X1187638033','Ninette','Gandey'),(30,'Y1800815354','Illa','Weeke'),(31,'E1937355288','Howard','Kidney'),(32,'I1679749547','Erina','Cory'),(33,'V1573946143','Amos','Hare'),(34,'T1337018417','Toma','Ransley'),(35,'Z1764114272','Mirabella','Hinchon'),(36,'I1784996824','Gracia','Battram'),(37,'P1542620848','Hagen','Foy'),(38,'Q1751605181','Kerwin','Pren'),(39,'I1870825125','Haleigh','Chellenham'),(40,'F1924380289','Charo','Pessler'),(41,'X1903943167','Piper','Arnecke'),(42,'M1066271348','Addie','Thies'),(43,'Z1791416239','Garvy','Yegorev'),(44,'Y1823539219','Gussy','Domnick'),(45,'R1043272038','Jayne','Pomfret'),(46,'Z1821791067','Talia','Bortolini'),(47,'Z1853276075','Dre','Crossan'),(48,'O1963451448','Hillery','Shapiro'),(49,'F1310005117','Kathleen','Sommerlin'),(50,'P1553876757','Jonis','Bowdrey'),(51,'A1056501628','Jacki','Faircliff'),(52,'X1957661861','Cindee','Constant'),(53,'E1213211617','Lottie','Jozwicki'),(54,'W1734915011','Antonella','D''Agostino'),(55,'T1318662379','Court','Balnaves'),(56,'T1961272073','Norene','Clover'),(57,'E1607157598','Heidi','Safe'),(58,'L1732713137','Raynard','Frantz'),(59,'B1369421289','Niel','Speer'),(60,'X1589996786','Ed','Adamolli'),(61,'N1602254291','Chaunce','Eustis'),(62,'F1575745471','Hugibert','Frankland'),(63,'G1375494771','Camila','Gayden'),(64,'K1593212814','Errol','Vousden'),(65,'M1585428408','Lezley','Theodoris'),(66,'Z1960616057','Anni','Brattan'),(67,'Y1872933671','Marthena','Blackway'),(68,'N1327215348','Emmy','Ovill'),(69,'W1015397499','Kamillah','Drivers'),(70,'V1023674882','Ivie','Tejero'),(71,'N1849354309','Winnah','Bereford'),(72,'H1632172317','Trenton','Caddens'),(73,'P1069015721','Melicent','Jouannin'),(74,'N1533973463','Hillary','Barrus'),(75,'N1759108960','Marrissa','Tarren'),(76,'V1479428879','Hallsy','Hexham'),(77,'C1885194322','Janna','Laneham'),(78,'E1805791323','Davidde','Millican'),(79,'H1997317036','Dane','Hearson'),(80,'L1750614103','Normie','Will'),(81,'G1988051787','Debora','Reeson'),(82,'D1206329133','Minnie','Giacomello'),(83,'H1873951165','Brandais','Paulino'),(84,'H1874207613','Nevin','Marciek'),(85,'S1118966327','Burch','Fursey'),(86,'N1247071431','Toma','Arrigucci'),(87,'N1519691826','Carrissa','Swalteridge'),(88,'J1684370878','Stanly','McGennis'),(89,'X1002971089','Joane','Bernardy'),(90,'B1851934462','Beilul','Weine'),(91,'G1801002867','Berni','Itschakov'),(92,'W1591913151','Chadd','Stainsby'),(93,'A1077412843','Christie','Pickrell'),(94,'Z1740499255','Philly','Elderbrant'),(95,'G1497859760','Saul','Higgonet'),(96,'S1180811648','Amery','Casson'),(97,'T1513422716','Marmaduke','Opdenort'),(98,'P1794039210','Dion','Coburn'),(99,'Y1377056753','Cordell','Belbin'),(100,'R1986140910','Kareem','Tollet');
/*!40000 ALTER TABLE `buyers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:10:07
