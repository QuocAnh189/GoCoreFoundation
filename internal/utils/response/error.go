package response

import (
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

type AppError struct {
	Message string
	Debug   string
	Status  status.AppStatusCode
}

func (e *AppError) Error() string {
	return fmt.Sprintf("MError: %s", e.Message)
}
