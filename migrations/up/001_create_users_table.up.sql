-- migration up
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(128) NOT NULL,
  `middle_name` varchar(128) DEFAULT NULL,
  `last_name` varchar(128) NOT NULL,
  `phone` varchar(128) NOT NULL,
  `email` varchar(128) NOT NULL,
  `role` varchar(16) NOT NULL,
  `status` varchar(16) NOT NULL,
  `create_id` INT DEFAULT 0,
  `create_dt` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `modify_id` INT DEFAULT 0,
  `modify_dt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO users (
    first_name, middle_name, last_name, phone, email, role, status
) VALUES
    ('Alice', NULL, 'Smith', '1234567890', 'alice.smith@example.com', 'admin', 'ACTIVE'),
    ('Bob', 'James', 'Johnson', '2345678901', 'bob.johnson@example.com', 'user', 'ACTIVE'),
    ('Carol', NULL, 'Williams', '3456789012', 'carol.williams@example.com', 'user', 'ACTIVE'),
    ('David', 'Michael', 'Brown', '4567890123', 'david.brown@example.com', 'guest', 'ACTIVE'),
    ('Eve', NULL, 'Davis', '5678901234', 'eve.davis@example.com', 'user', 'ACTIVE');