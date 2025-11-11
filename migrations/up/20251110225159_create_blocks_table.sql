CREATE TABLE `blocks` (
  `id` char(36) NOT NULL,
  `type` varchar(255) NOT NULL, -- DEVICE, EMAIL, PHONE, IP
  `value` varchar(255) NOT NULL,
  `reason` varchar(255) DEFAULT NULL, -- send max exceed otp
  `blocked_dt` datetime(3) DEFAULT NULL,
  `blocked_until_dt` datetime(3) DEFAULT NULL,
  `status` varchar(32) DEFAULT NULL,  -- ACTIVE or ACTIVE
  `create_id` int DEFAULT '0',
  `create_dt` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `modify_id` int DEFAULT '0',
  `modify_dt` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_dt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=242 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci