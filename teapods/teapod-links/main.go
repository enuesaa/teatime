package main

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/go-playground/validator/v10"
)

func main() {
	plug.Serve(&Provider{})
}

type Provider struct {}

func (p *Provider) Info() (plug.Info, error) {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
		Teaboxes: []plug.Teabox{
			{
				Name: "links",
			},
		},
		Actions: []plug.Action{
			{
				Name:   "action.sync",
				Comment: "sync links",
			},
		},
	}
	return info, nil
}

func (p *Provider) On(event plug.Event) (plug.Logs, error) {
	logs := plug.NewLogs()
	logs.Info("app start")

	if event.Name == "data.created" {
		errs := validator.New().ValidateMap(event.Data, rules)
		for name, value := range errs {
			return logs, fmt.Errorf("%s: %s", name, value)
		}
		return logs, nil
	}

	return logs, nil
}
