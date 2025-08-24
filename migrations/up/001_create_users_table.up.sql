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
  `create_id` int DEFAULT NULL,
  `create_dt` timestamp NULL DEFAULT NULL,
  `modify_id` int DEFAULT NULL,
  `modify_dt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO users (
    first_name, middle_name, last_name, phone, email, role, status, create_id, create_dt, modify_id, modify_dt
) VALUES
    ('Alice', NULL, 'Smith', '1234567890', 'alice.smith@example.com', 'admin', 'active', 1, CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP),
    ('Bob', 'James', 'Johnson', '2345678901', 'bob.johnson@example.com', 'user', 'active', 1, CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP),
    ('Carol', NULL, 'Williams', '3456789012', 'carol.williams@example.com', 'user', 'active', 1, CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP),
    ('David', 'Michael', 'Brown', '4567890123', 'david.brown@example.com', 'guest', 'active', NULL, CURRENT_TIMESTAMP, NULL, CURRENT_TIMESTAMP),
    ('Eve', NULL, 'Davis', '5678901234', 'eve.davis@example.com', 'user', 'active', NULL, CURRENT_TIMESTAMP, NULL, CURRENT_TIMESTAMP);