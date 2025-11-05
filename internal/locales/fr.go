package locales

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"

var (
	FR LanguageType = "fr"
)

func GetMessageFRFromStatus(statusCode status.Code) string {
	switch statusCode {
	case status.USER_INVALID_PARAMS:
		return "Paramètres invalides"
	case status.USER_INVALID_ID:
		return "ID utilisateur invalide"
	case status.USER_NOT_FOUND:
		return "Utilisateur non trouvé"
	case status.USER_MISSING_FIRST_NAME:
		return "Le prénom est requis"
	case status.USER_MISSING_LAST_NAME:
		return "Le nom de famille est requis"
	case status.USER_MISSING_EMAIL:
		return "L'email est requis"
	case status.USER_INVALID_EMAIL:
		return "Format d'email invalide"
	case status.USER_MISSING_PHONE:
		return "Le téléphone est requis"
	case status.USER_INVALID_ROLE:
		return "Rôle invalide"
	case status.USER_INVALID_STATUS:
		return "Statut invalide"
	case status.SUCCESS:
		return "Succès"
	default:
		return "Unknown error"
	}
}
