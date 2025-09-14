CREATE TABLE lingos (
  `id` CHAR(36) NOT NULL,
  `lang` varchar(8) NOT NULL,
  `key` varchar(64) NOT NULL,
  `val` text NOT NULL,
  `status` varchar(16) NOT NULL,
  `create_dt` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `create_id` INT DEFAULT 0,
  `modify_dt` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `modify_id` INT DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO lingos (
  `id`, `lang`, `key`, `val`, `status`
) VALUES
  ('018a7b3e-4b3e-7b3e-8b3e-4b3e7b3e4b3e', 'en', 'common_error_invalid_params', 'Invalid params', 'ACTIVE'),
  ('018a7b3f-4b3f-7b3f-8b3f-4b3f7b3f4b3f', 'vn', 'common_error_invalid_params', 'Thông tin yêu cầu không đúng', 'ACTIVE'),
  ('018a7b40-4b40-7b40-8b40-4b407b404b40', 'en', 'common_error_internal', 'Something went wentwrong, please try it later', 'ACTIVE'),
  ('018a7b41-4b41-7b41-8b41-4b417b414b41', 'vn', 'common_error_internal', 'Đang xảy ra lỗi, vui lòng thử lại sau', 'ACTIVE');