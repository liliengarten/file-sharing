package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Errors []ValidationError `json:"errors"`
}

func Validate[T any](object T) *ErrorResponse {
	err := validator.New().Struct(object)

	if err != nil {
		var errs validator.ValidationErrors

		if errors.As(err, &errs) {
			var formattedErrors []ValidationError

			for _, e := range errs {
				formattedErrors = append(formattedErrors, ValidationError {
					Field: e.Field(),
					Message: e.Tag(),
				})
			}

			resp := ErrorResponse{
				Message: "Validation error",
				Errors: formattedErrors,
			}

			return &resp
		}
	}

	return &ErrorResponse{}
}
