/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.5.21 : Database - dbdatabase
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`dbdatabase` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `dbdatabase`;

/*Table structure for table `dbmaccount` */

DROP TABLE IF EXISTS `dbmaccount`;

CREATE TABLE `dbmaccount` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `Username` varchar(30) NOT NULL,
  `Password` varchar(30) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

/*Data for the table `dbmaccount` */

insert  into `dbmaccount`(`ID`,`Username`,`Password`) values 
(7,'fffsdfsdf','fffsdfsd'),
(8,'Icshan','Icshan'),
(9,'adfhadg','asdgasd'),
(10,'sfdsfff','sdfsdfff');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
