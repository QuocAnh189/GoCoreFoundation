package locales

import (
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

var (
	VN LanguageType = "vn"
)

func GetMessageVNFromStatus(statusCode status.Code, args ...any) string {
	switch statusCode {
	case status.USER_INVALID_PARAMS:
		return "Tham số không hợp lệ"
	case status.USER_INVALID_ID:
		return "ID người dùng không hợp lệ"
	case status.USER_NOT_FOUND:
		return "Không tìm thấy người dùng"
	case status.USER_MISSING_FIRST_NAME:
		return "Thiếu tên"
	case status.USER_MISSING_LAST_NAME:
		return "Thiếu họ"
	case status.USER_MISSING_EMAIL:
		return "Thiếu email"
	case status.USER_INVALID_EMAIL:
		return "Định dạng email không hợp lệ"
	case status.USER_EMAIL_ALREADY_EXISTS:
		return "Email đã tồn tại"
	case status.USER_MISSING_PHONE:
		return "Thiếu số điện thoại"
	case status.USER_INVALID_PHONE:
		return "Định dạng số điện thoại không hợp lệ"
	case status.USER_PHONE_ALREADY_EXISTS:
		return "Số điện thoại đã tồn tại"
	case status.USER_INVALID_ROLE:
		return fmt.Sprintf("Vai trò không hợp lệ. Các vai trò hợp lệ là: %v", args)
	case status.USER_INVALID_STATUS:
		return fmt.Sprintf("Trạng thái không hợp lệ. Các trạng thái hợp lệ là: %v", args)
	case status.SUCCESS:
		return "Thành công"
	default:
		return "Unknown"
	}
}
