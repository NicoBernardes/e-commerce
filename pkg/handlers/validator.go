package handlers

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)

	validationErrors := err.(validator.ValidationErrors)
	validateError := validationErrors[0]

	field := strings.ToLower(validateError.StructField())

	// Can upgrade with more cases
	switch validateError.Tag() {
	case "required":
		return errors.New(field + " is required")
	}
	return nil
}
