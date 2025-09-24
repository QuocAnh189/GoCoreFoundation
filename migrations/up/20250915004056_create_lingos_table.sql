CREATE TABLE lingos (
  `id` CHAR(36) NOT NULL,
  `lang` varchar(8) NOT NULL,
  `lkey` varchar(64) NOT NULL,
  `lval` text NOT NULL,
  `status` VARCHAR(16) NOT NULL DEFAULT 'ACTIVE',
  `create_id` INT DEFAULT 0,
  `create_dt` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `modify_id` INT DEFAULT 0,
  `modify_dt` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_dt` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO lingos (
  `id`, `lang`, `lkey`, `lval`, `status`
) VALUES
  ('018a7b3e-4b3e-7b3e-8b3e-4b3e7b3e4b3e', 'en', 'common_error_invalid_params', 'Invalid params', 'ACTIVE'),
  ('018a7b3f-4b3f-7b3f-8b3f-4b3f7b3f4b3f', 'vn', 'common_error_invalid_params', 'Thông tin yêu cầu không đúng', 'ACTIVE'),
  ('018a7b40-4b40-7b40-8b40-4b407b404b40', 'en', 'common_error_internal', 'Something went wentwrong, please try it later', 'ACTIVE'),
  ('018a7b41-4b41-7b41-8b41-4b417b414b41', 'vn', 'common_error_internal', 'Đang xảy ra lỗi, vui lòng thử lại sau', 'ACTIVE'),
  -- user.invalid_parameter
  ('018b8c4a-5c4a-7c4a-8c4a-5c4a7c4a5c4a', 'en', 'user.invalid_parameter', 'Invalid parameters', 'ACTIVE'),
  ('018b8c4b-5c4b-7c4b-8c4b-5c4b7c4b5c4b', 'vn', 'user.invalid_parameter', 'Thông số không hợp lệ', 'ACTIVE'),
  -- user.invalid_user_id
  ('018b8c4c-5c4c-7c4c-8c4c-5c4c7c4c5c4c', 'en', 'user.invalid_user_id', 'Invalid user ID', 'ACTIVE'),
  ('018b8c4d-5c4d-7c4d-8c4d-5c4d7c4d5c4d', 'vn', 'user.invalid_user_id', 'ID người dùng không hợp lệ', 'ACTIVE'),
  -- user.not_found
  ('018b8c4e-5c4e-7c4e-8c4e-5c4e7c4e5c4e', 'en', 'user.not_found', 'User not found', 'ACTIVE'),
  ('018b8c4f-5c4f-7c4f-8c4f-5c4f7c4f5c4f', 'vn', 'user.not_found', 'Không tìm thấy người dùng', 'ACTIVE'),
  -- user.first_name_required
  ('018b8c50-5c50-7c50-8c50-5c507c505c50', 'en', 'user.first_name_required', 'First name is required', 'ACTIVE'),
  ('018b8c51-5c51-7c51-8c51-5c517c515c51', 'vn', 'user.first_name_required', 'Tên là bắt buộc', 'ACTIVE'),
  -- usre_.last_name_required
  ('018b8c52-5c52-7c52-8c52-5c527c525c52', 'en', 'usre.last_name_required', 'Last name is required', 'ACTIVE'),
  ('018b8c53-5c53-7c53-8c53-5c537c535c53', 'vn', 'usre.last_name_required', 'Họ là bắt buộc', 'ACTIVE'),
  -- user.phone_required
  ('018b8c54-5c54-7c54-8c54-5c547c545c54', 'en', 'user.phone_required', 'Phone number is required', 'ACTIVE'),
  ('018b8c55-5c55-7c55-8c55-5c557c555c55', 'vn', 'user.phone_required', 'Số điện thoại là bắt buộc', 'ACTIVE'),
  -- user.email_required
  ('018b8c56-5c56-7c56-8c56-5c567c565c56', 'en', 'user.email_required', 'Email is required', 'ACTIVE'),
  ('018b8c57-5c57-7c57-8c57-5c577c575c57', 'vn', 'user.email_required', 'Email là bắt buộc', 'ACTIVE'),
  -- user.invalid_email_format
  ('018b8c58-5c58-7c58-8c58-5c587c585c58', 'en', 'user.invalid_email_format', 'Invalid email format', 'ACTIVE'),
  ('018b8c59-5c59-7c59-8c59-5c597c595c59', 'vn', 'user.invalid_email_format', 'Định dạng email không hợp lệ', 'ACTIVE'),
  -- user.invalid_role
  ('018b8c5a-5c5a-7c5a-8c5a-5c5a7c5a5c5a', 'en', 'user.invalid_role', 'Invalid role', 'ACTIVE'),
  ('018b8c5b-5c5b-7c5b-8c5b-5c5b7c5b5c5b', 'vn', 'user.invalid_role', 'Vai trò không hợp lệ', 'ACTIVE'),
  -- user.invalid_status
  ('018b8c5c-5c5c-7c5c-8c5c-5c5c7c5c5c5c', 'en', 'user.invalid_status', 'Invalid status', 'ACTIVE'),
  ('018b8c5d-5c5d-7c5d-8c5d-5c5d7c5d5c5d', 'vn', 'user.invalid_status', 'Trạng thái không hợp lệ', 'ACTIVE');