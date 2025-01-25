package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateLinkTea(data map[string]interface{}) error {
	rule := map[string]interface{}{
		"title": "required",
		"link": "required,url",
	}

	errs := validator.New().ValidateMap(data, rule)

	if len(errs) > 0 {
		message := ""
		for key, value := range errs {
			message = fmt.Sprintf("\n  %s: %s", key, value)
		}
		return fmt.Errorf("validation failed: %s", message)
	}
	return nil
}
