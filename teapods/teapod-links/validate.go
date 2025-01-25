package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func fmtValidationErrs(errs map[string]interface{}) error {
	message := ""
	for field, value := range errs {
		fieldErrs := value.(validator.ValidationErrors)
		for _, err := range fieldErrs {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s\n  Required: %s", message, field)
			case "url":
				message = fmt.Sprintf("%s\n  InvalidUrl: %s", message, field)
			}
		}
	}
	return fmt.Errorf("validation failed: %s", message)
}

func ValidateLinkTea(data map[string]interface{}) error {
	rule := map[string]interface{}{
		"title": "required",
		"link": "required,url",
	}

	errs := validator.New().ValidateMap(data, rule)

	if len(errs) > 0 {
		return fmtValidationErrs(errs)
	}
	return nil
}

func ValidateNoteTea(data map[string]interface{}) error {
	rule := map[string]interface{}{
		"title": "required",
		"memo": "required",
	}

	errs := validator.New().ValidateMap(data, rule)

	if len(errs) > 0 {
		return fmtValidationErrs(errs)
	}
	return nil
}
