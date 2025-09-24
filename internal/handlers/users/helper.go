package users

import (
	locale_en "github.com/QuocAnh189/GoCoreFoundation/internal/constants/locale/en"
	locale_fr "github.com/QuocAnh189/GoCoreFoundation/internal/constants/locale/fr"
	locale_vn "github.com/QuocAnh189/GoCoreFoundation/internal/constants/locale/vn"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

func DetermineErrKey(err error) string {
	switch err {
	case ErrInvalidParameter:
		return "user.invalid_parameter"
	case ErrInvalidUserID:
		return "user.invalid_user_id"
	case ErrUserNotFound:
		return "user.not_found"
	case ErrMissingFirstName:
		return "user.first_name_required"
	case ErrMissingLastName:
		return "user.last_name_required"
	case ErrMissingPhone:
		return "user.phone_required"
	case ErrMissingEmail:
		return "user.email_required"
	case ErrInvalidEmail:
		return "user.invalid_email_format"
	case ErrInvalidRole:
		return "user.invalid_role"
	case ErrInvalidStatus:
		return "user.invalid_status"
	default:
		return "user.unknown_error"
	}
}

func DetermineErrStatus(err error) int {
	switch err {
	case ErrInvalidParameter, ErrInvalidUserID, ErrMissingFirstName,
		ErrMissingLastName, ErrMissingPhone, ErrMissingEmail,
		ErrInvalidEmail, ErrInvalidRole, ErrInvalidStatus:
		return status.BAD_REQUEST
	case ErrUserNotFound:
		return status.NOT_FOUND
	default:
		return status.INTERNAL
	}
}

func GetMessageFromKey(lang, key string) string {
	switch lang {
	case locale_vn.LocaleLanguageVN:
		if val, ok := locale_vn.GetMessageFromKeyVN[key]; ok {
			return val
		}
	case locale_en.LocaleLanguageEN:
		if val, ok := locale_en.GetMessageFromKeyEN[key]; ok {
			return val
		}
	case locale_fr.LocaleLanguageFR:
		if val, ok := locale_fr.GetMessageFromKeyFR[key]; ok {
			return val
		}
	default:
		if val, ok := locale_en.GetMessageFromKeyEN[key]; ok {
			return val
		}
	}

	return "Unknown error"
}
