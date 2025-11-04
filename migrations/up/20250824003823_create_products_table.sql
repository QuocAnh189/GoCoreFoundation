-- migration up
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `description` varchar(256) DEFAULT NULL,
  `price` decimal(10,2) NOT NULL,
  `create_id` int DEFAULT NULL,
  `create_dt` timestamp NULL DEFAULT NULL,
  `modify_id` int DEFAULT NULL,
  `modify_dt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
