package error_handler

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ErrValidation = errors.New("validation_failed")

type FieldError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message,omitempty"`
}

type ValidationError struct {
	Fields []FieldError
}

func (e *ValidationError) Error() string { // satisfies error
	return fmt.Sprintf("%v: %d field error(s)", ErrValidation, len(e.Fields))
}

func NewValidationError(err error) *ValidationError {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}
	fields := make([]FieldError, 0, len(ve))
	for _, fe := range ve {
		fields = append(fields, FieldError{
			Field: fe.Field(), // struct field (you can map to JSON name if you prefer)
			Tag:   fe.Tag(),   // e.g. "required", "email", "min"
			Param: fe.Param(), // e.g. "3" for min=3
			// Optional: customize Message per tag if you want user-friendly text
			Message: fmt.Sprintf("%s failed on '%s'", fe.Field(), fe.Tag()),
		})
	}
	return &ValidationError{Fields: fields}
}
