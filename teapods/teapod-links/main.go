package main

import (
	"encoding/json"
	"log"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func main() {
	provider := Provider{
		Repos: repository.New(),
	}
	if err := provider.Repos.DB.Open(); err != nil {
		log.Fatalf("Error: %s", err)
	}
	plug.Serve(&provider)
	if err := provider.Repos.DB.Close(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type Provider struct {
	Repos repository.Repos
}

func (p *Provider) Info() (plug.Info, error) {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
		Cards: []string{},
		Schemas: []plug.Schema{
			{
				Name: "links",
				Vals: map[string]string{
					"title": "str",
					"link": "str",
				},
			},
		},
	}
	return info, nil
}

func (p *Provider) List() ([]string, error) {
	dbteas, err := p.Repos.DB.ListTeas("links")
	if err != nil {
		return make([]string, 0), err
	}
	list := make([]string, 0)
	for _, dbtea := range dbteas {
		list = append(list, dbtea.Teaid)
	}
	return list, nil
}

func (p *Provider) Get(teaid string) (plug.Tea, error) {
	dbtea, err := p.Repos.DB.GetTea("links", teaid)
	if err != nil {
		return plug.Tea{}, err
	}
	var value plug.Value
	if err := json.Unmarshal([]byte(dbtea.Value.(string)), &value); err != nil {
		return plug.Tea{}, err
	}
	return plug.Tea{Teaid: dbtea.Teaid, Value: value}, nil
}

func (p *Provider) Set(tea plug.Tea) error {
	valuebytes, err := json.Marshal(tea.Value)
	if err != nil {
		return err
	}
	if err := p.Repos.DB.CreateTea("links", tea.Teaid, string(valuebytes)); err != nil {
		return err
	}
	return nil
}

func (p *Provider) Del(teaid string) error {
	if err := p.Repos.DB.DeleteTea("links", teaid); err != nil {
		return err
	}
	return nil
}

func (p *Provider) GetCard(name string) (plug.Card, error) {
	return plug.Card{}, nil
}
