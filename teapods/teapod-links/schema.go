package main

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type LinkTea struct {
	Title string `validate:"required"`
	Link string `validate:"required,url"`
}

func ValidateLinkTea(data []byte) error {
	var tea LinkTea
	if err := json.Unmarshal(data, &tea); err != nil {
		return err
	}
	if err := validator.New().Struct(tea); err != nil {
		return err
	}
	return nil
}
