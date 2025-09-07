package response

import (
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

type AppError struct {
	Message   string
	Status    status.AppStatusCode
	BaseError error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error: %s", e.BaseError.Error())
}

func ErrInvalidParams() *AppError {
	return &AppError{
		Message:   "Invalid Params",
		Status:    status.BAD_REQUEST,
		BaseError: fmt.Errorf("invalid parameters"),
	}
}

func ErrInvalidParseString(value string) *AppError {
	return &AppError{
		Message:   "Invalid parse ID",
		Status:    status.BAD_REQUEST,
		BaseError: fmt.Errorf("invalid parse ID from string: %s", value),
	}
}
