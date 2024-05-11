package validation

import "github.com/go-playground/validator/v10"

type CustomValidationError struct {
	HasError bool
	Field    string
	Tag      string
	Param    string
	Value    interface{}
}

var validate = validator.New()

func Validate(data interface{}) []CustomValidationError {
	var customValidationError []CustomValidationError

	if errors := validate.Struct(data); errors != nil {
		for _, fieldError := range errors.(validator.ValidationErrors) {
			var cve CustomValidationError
			cve.HasError = true
			cve.Field = fieldError.Field()
			cve.Tag = fieldError.Tag()
			cve.Param = fieldError.Param()
			cve.Value = fieldError.Value()
			customValidationError = append(customValidationError, cve)
		}
	}

	return customValidationError
}
