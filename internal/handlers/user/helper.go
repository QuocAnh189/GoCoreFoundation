package user

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

func GetArgsByStatatus(statusCode status.Code) []any {
	switch statusCode {
	case status.USER_INVALID_ROLE:
		return []any{enum.RoleGuest, enum.RoleUser, enum.RoleAdmin}
	case status.USER_INVALID_STATUS:
		return []any{enum.StatusActive, enum.StatusInactive, enum.StatusBanned}
	default:
		return []any{}
	}
}
