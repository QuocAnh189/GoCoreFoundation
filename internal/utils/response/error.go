package response

import (
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants"
)

type AppError struct {
	Message string
	Status  constants.AppStatusCode
}

func (e *AppError) Error() string {
	return fmt.Sprintf("MError: %s", e.Message)
}
