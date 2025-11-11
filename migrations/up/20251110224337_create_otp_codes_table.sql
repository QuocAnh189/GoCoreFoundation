CREATE TABLE `otp_codes` (
  `id` char(36) NOT NULL,
  `purpose` varchar(32) DEFAULT NULL, -- 2FA, SIGNUP, FORGOT_PASSWORD
  `uid` char(36) DEFAULT NULL,
  `identifier` varchar(255) DEFAULT NULL, 
  `device_uuid` varchar(255) DEFAULT NULL,
  `device_name` varchar(255) DEFAULT NULL,
  `gen_otp_cnt` int NOT NULL DEFAULT '0',
  `verify_otp_cnt` int NOT NULL DEFAULT '0',
  `otp_code` varchar(32) NOT NULL,
  `otp_create_dt` datetime(3) DEFAULT NULL,
  `otp_expire_dt` datetime(3) DEFAULT NULL,
  `status` varchar(16) NOT NULL DEFAULT 'ACTIVE',
  `create_id` int DEFAULT '0',
  `create_dt` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `modify_id` int DEFAULT '0',
  `modify_dt` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_dt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_uid_identifier_device` (`uid`,`identifier`,`device_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=3245 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci