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
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `expiration_rate` int NOT NULL,
  `freezing_rate` int NOT NULL,
  `height` float NOT NULL,
  `length` float NOT NULL,
  `netweight` float NOT NULL,
  `product_code` varchar(10) COLLATE utf8mb4_general_ci NOT NULL,
  `recommended_freezing_temperature` int NOT NULL,
  `width` int NOT NULL,
  `product_type_id` int NOT NULL,
  `seller_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_code` (`product_code`),
  KEY `product_type_id` (`product_type_id`),
  KEY `seller_id` (`seller_id`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_type` (`id`),
  CONSTRAINT `products_ibfk_2` FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Leche entera 1L',7,1,25,8,1,'AO6A9H',-18,8,1,9),(2,'Queso fresco 200g',10,1,6,10,0.2,'YYSGOM',-18,6,1,61),(3,'Vestibulum ante ipsum...',5,9,83.5,87.2,47.7,'YYSGOMM',0,1,14,12),(4,'Leche descremada 1L',7,1,25,8,1,'AI0M96',-18,8,1,69),(5,'Crema de leche 250ml',15,1,15,7,0.25,'6LR112',-18,7,1,20),(6,'Manzanas Golden 1kg',14,0,10,20,1,'9TIDTQ',0,20,2,42),(7,'Tomates 500g',7,0,10,10,0.5,'VPWVF4',0,10,2,36),(8,'Zanahorias 1kg',14,0,15,15,1,'0B16SI',0,15,2,100),(9,'Lechuga 300g',7,0,10,10,0.3,'PKHEZ7',0,10,2,5),(10,'Papas 1kg',30,0,20,10,1,'52ET5G',0,10,2,4),(11,'Carne de res molida 500g',3,2,5,15,0.5,'DYRNBC',-18,10,3,36),(12,'Vestibulum ante ipsum...',5,9,6,87.2,47.7,'5D3IDT',-18,1,14,12),(13,'Pechuga de pavo 500g',7,2,6,20,0.5,'DCA5N3',-18,10,3,96),(14,'Filetes de salmón 300g',5,2,4,15,0.3,'WUFKV4',-18,10,3,36),(15,'Costillas de cerdo 1kg',5,2,6,20,1,'DSQSDI',-18,10,3,78),(16,'Pechuga de pollo congelada 1kg',7,2,6,20,1,'C4XQKG',-18,10,4,31),(17,'Pizza congelada 300g',30,1,5,15,0.3,'0Q9QB1',-18,15,4,66),(18,'Verduras congeladas 500g',90,1,10,20,0.5,'STAK4D',-18,10,4,3),(19,'Helado de vainilla 500g',180,1,10,8,0.5,'CZMJ0B',-18,8,4,75),(20,'Hamburguesas congeladas 4 piezas',60,2,5,20,0.4,'8L3J1N',-18,10,4,77),(21,'Jugo de naranja 1L',7,0,25,8,1,'QZRHJ8',0,8,5,11),(22,'Cerveza 330ml',180,0,20,7,0.33,'7JE86I',0,7,5,24),(23,'Refresco de cola 2L',180,0,30,10,2,'VOZZZG',0,10,5,90),(24,'Agua mineral 500ml',365,0,10,5,0.5,'ACF243',0,5,5,81),(25,'Té verde 1L',14,0,25,8,1,'70HT4X',0,8,5,92),(26,'Arroz blanco 1kg',365,0,20,10,1,'6GM5JA',0,10,6,66),(27,'Pasta de trigo 500g',365,0,15,7,0.5,'IU576K',0,7,6,28),(28,'Fideos 500g',365,0,15,7,0.5,'A2YYAK',0,7,6,45),(29,'Quinua 500g',365,0,15,7,0.5,'QJYL6J',0,7,6,69),(30,'Avena 500g',365,0,15,7,0.5,'LN7KM9',0,7,6,44),(31,'Pan de caja 500g',7,0,10,15,0.5,'EHW0R9',-18,10,7,61),(32,'Galletas de avena 300g',30,0,5,20,0.3,'DTIERF',0,10,7,3),(33,'Pastel de chocolate 250g',14,0,6,15,0.25,'75D1UZ',0,15,7,47),(34,'Cereal de maíz 300g',90,0,8,10,0.3,'DVV81P',0,8,7,14),(35,'Galletas saladas 200g',60,0,4,10,0.2,'WI1B13',0,10,7,64),(36,'Papel higiénico 12 rollos',365,0,25,10,1,'GJ0ACS',0,10,8,62),(37,'Jabón de baño 3 piezas',365,0,5,10,0.3,'NS4K9D',0,5,8,19),(38,'Detergente líquido 1L',365,0,25,10,1,'S0ARZ7',0,10,8,3),(39,'Pasta dental 150g',365,0,10,5,0.15,'NIFDJZ',0,5,8,42),(40,'Shampoo 400ml',365,0,15,7,0.4,'SEPKOL',0,7,8,17),(41,'Aceite de oliva 500ml',365,0,20,8,0.5,'6PRBBZ',0,8,9,22),(42,'Vinagre de manzana 500ml',365,0,20,8,0.5,'39MN65',0,8,9,22),(43,'Sal de cocina 1kg',365,0,20,10,1,'VR22R5',0,10,9,72),(44,'Azúcar blanca 1kg',365,0,20,10,1,'Z31LSV',0,10,9,37),(45,'Mostaza 250g',365,0,15,7,0.25,'IXU4O9',0,7,9,89),(46,'Atún enlatado 200g',365,0,8,8,0.2,'IX3VTD',0,8,10,74),(47,'Sopa enlatada 500g',365,0,12,10,0.5,'2ZQUKQ',0,10,10,94),(48,'Frijoles enlatados 400g',365,0,10,8,0.4,'2WF405',0,8,10,73),(49,'Tomates enlatados 400g',365,0,10,8,0.4,'QKNSP9',0,8,10,49),(50,'Piña en almibar 400g',365,0,10,8,0.4,'94JMR4',0,8,10,93),(51,'Limpiador multiuso 1L',365,0,25,10,1,'GSGIM7',0,10,11,8),(52,'Desinfectante líquido 500ml',365,0,20,8,0.5,'ATQ1VK',0,8,11,39),(53,'Lavavajillas líquido 750ml',365,0,20,8,0.75,'DN937E',0,8,11,83),(54,'Detergente en polvo 2kg',365,0,25,15,2,'TFL6WU',0,10,11,56),(55,'Limpiador de pisos 1L',365,0,25,10,1,'GCLTMJ',0,10,11,68),(56,'Martillo 500g',0,0,35,15,0.5,'CX8DY9',0,5,12,50),(57,'Destornillador 6 piezas',0,0,25,5,0.3,'SRX4X0',0,5,12,81),(58,'Cinta métrica 5m',0,0,5,15,0.2,'FJPEEY',0,5,12,30),(59,'Alicates 200mm',0,0,25,7,0.4,'HJA1SN',0,5,12,19),(60,'Llave inglesa 250mm',0,0,30,10,0.5,'NQHSTC',0,5,12,47),(61,'Camisa de algodón tamaño M',365,0,5,30,0.3,'2UZFN2',0,20,13,91),(62,'Pantalones de mezclilla tamaño 32',365,0,5,40,0.8,'44DXLJ',0,20,13,78),(63,'Zapatillas deportivas talla 42',365,0,12,30,0.9,'RU1CYD',0,20,13,73),(64,'Camiseta deportiva talla L',365,0,5,30,0.2,'IMDBP9',0,20,13,54),(65,'Botas de lluvia talla 40',365,0,25,30,1.2,'E60YNO',0,25,13,88),(66,'Muñeca de trapo',365,0,30,15,0.4,'UG6OD9',0,10,14,42),(67,'Lego set 500 piezas',365,0,20,15,1.5,'Y16GMA',0,15,14,84),(68,'Pelota de fútbol tamaño 5',365,0,20,20,0.5,'QT72LO',0,20,14,76),(69,'Pista de autos eléctrica',365,0,10,30,1.5,'P5FYVS',0,20,14,84),(70,'Juego de mesa Monopoly',365,0,5,20,1,'ARGSDQ',0,20,14,61);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-15 11:10:14
