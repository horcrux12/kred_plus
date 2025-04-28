package exception

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Message  string
	Code     int
	HttpCode int
}

func (c CustomError) Error() string {
	return c.Message
}

var (
	InternalServerError = CustomError{
		Message:  "Internal Server Error, Please contact our customer service",
		Code:     1000,
		HttpCode: http.StatusInternalServerError,
	}

	NotFoundError = func(field string) CustomError {
		return CustomError{
			Message:  fmt.Sprintf("no data found with this '%s'", field),
			Code:     1001,
			HttpCode: http.StatusNotFound,
		}
	}

	ForbiddenAccess = CustomError{
		Message:  "You doesn't have access",
		Code:     1002,
		HttpCode: http.StatusForbidden,
	}

	InvalidRequest = CustomError{
		Message:  "Invalid request",
		Code:     1003,
		HttpCode: http.StatusUnprocessableEntity,
	}

	InvalidRequestWithMessage = func(field, detail string) error {
		message := fmt.Sprintf("Invalid request for this %s", field)
		if detail != "" {
			message = fmt.Sprintf("%s, %s", message, detail)
		}
		return CustomError{
			Message:  message,
			Code:     1004,
			HttpCode: http.StatusUnprocessableEntity,
		}
	}

	AlreadyExist = func(field string) CustomError {
		return CustomError{
			Message:  fmt.Sprintf("These '%s' data already exist", field),
			Code:     1005,
			HttpCode: http.StatusConflict,
		}
	}

	MandatoryError = func(field string) CustomError {
		return CustomError{
			Message:  fmt.Sprintf("'%s' can't be empty", field),
			Code:     1006,
			HttpCode: http.StatusBadRequest,
		}
	}

	AlreadyExistWithMessage = func(field, message string) CustomError {
		return CustomError{
			Message:  fmt.Sprintf("These '%s' data already exist, %s", field, message),
			Code:     1005,
			HttpCode: http.StatusConflict,
		}
	}
)
