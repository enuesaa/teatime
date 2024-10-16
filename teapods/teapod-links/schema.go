package main

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type LinkTea struct {
	Link string `validate:"required,url"`
}

func BindLinkTea(data map[string]interface{}) (LinkTea, error) {
	var tea LinkTea
	jbytes, err := json.Marshal(data)
	if err != nil {
		return tea, err
	}
	if err := json.Unmarshal(jbytes, &tea); err != nil {
		return tea, err
	}
	if err := validator.New().Struct(tea); err != nil {
		return tea, err
	}
	return tea, nil
}
