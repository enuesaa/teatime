package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/plug"
)

func main() {
	db := plug.DB{}
	if err := db.Init("links"); err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer db.Close()

	provider := Provider{
		DB: db,
	}
	plug.Serve(&provider)
}

type Provider struct {
	DB plug.DB
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
	teas, err := p.DB.ListTeas()
	if err != nil {
		return make([]string, 0), err
	}
	list := make([]string, 0)
	for _, tea := range teas {
		list = append(list, tea.Teaid)
	}
	return list, nil
}

func (p *Provider) Get(teaid string) (plug.Tea, error) {
	return p.DB.GetTea(teaid)
}

func (p *Provider) Set(tea plug.Tea) error {
	return p.DB.CreateTea(tea)
}

func (p *Provider) Del(teaid string) error {
	return p.DB.DeleteTea(teaid)
}

func (p *Provider) GetCard(name string) (plug.Card, error) {
	return plug.Card{}, nil
}
