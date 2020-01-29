-- --------------------------------------------------------
-- Хост:                         127.0.0.1
-- Версия сервера:               8.0.15 - MySQL Community Server - GPL
-- Операционная система:         Win64
-- HeidiSQL Версия:              10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Дамп структуры базы данных kaspilab
CREATE DATABASE IF NOT EXISTS `kaspilab` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `kaspilab`;

-- Дамп структуры для таблица kaspilab.directions
CREATE TABLE IF NOT EXISTS `directions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `direction` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- Дамп данных таблицы kaspilab.directions: ~2 rows (приблизительно)
/*!40000 ALTER TABLE `directions` DISABLE KEYS */;
INSERT INTO `directions` (`id`, `direction`) VALUES
	(1, 'Data Science'),
	(2, 'Data Engineering'),
	(3, 'Мобильная разработка');
/*!40000 ALTER TABLE `directions` ENABLE KEYS */;

-- Дамп структуры для таблица kaspilab.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `fio` varchar(255) NOT NULL,
  `university` varchar(1000) DEFAULT NULL,
  `course` varchar(1000) DEFAULT NULL,
  `mobile_phone` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- Дамп данных таблицы kaspilab.users: ~17 rows (приблизительно)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- Дамп структуры для таблица kaspilab.user_direction
CREATE TABLE IF NOT EXISTS `user_direction` (
  `user_id` bigint(20) DEFAULT NULL,
  `direction_id` int(11) DEFAULT NULL,
  KEY `FK_user_direction_users` (`user_id`),
  KEY `FK_user_direction_directions` (`direction_id`),
  CONSTRAINT `FK_user_direction_directions` FOREIGN KEY (`direction_id`) REFERENCES `directions` (`id`),
  CONSTRAINT `FK_user_direction_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Дамп данных таблицы kaspilab.user_direction: ~0 rows (приблизительно)
/*!40000 ALTER TABLE `user_direction` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_direction` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
