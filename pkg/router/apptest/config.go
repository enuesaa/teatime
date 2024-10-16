package apptest

import (
	"io"
	"strings"
)

type Config struct {
	ParamNames  []string
	ParamValues []string
	Body        io.Reader
	Headers     map[string]string
}

func NewConfig() Config {
	return Config{
		ParamNames:  make([]string, 0),
		ParamValues: make([]string, 0),
		Body:        nil,
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},
	}
}

type UseFn func(*Config)

func UseParam(name string, value string) UseFn {
	return func(c *Config) {
		c.ParamNames = append(c.ParamNames, name)
		c.ParamValues = append(c.ParamValues, value)
	}
}

func UseBody(body string) UseFn {
	return func(c *Config) {
		c.Body = strings.NewReader(body)
	}
}

func UseHeader(name string, value string) UseFn {
	return func(c *Config) {
		c.Headers[name] = value
	}
}
