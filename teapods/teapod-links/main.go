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
		Cards: make([]string, 0),
		Teaboxes: []plug.Teabox{
			{
				Name: "links",
				Comment: "Resgister your favorite site and look up later.",
				Vals: map[string]string{
					"title": "str",
					"link": "str",
					"memo": "str",
					"priority": "str", // todo int
				},
			},
			{
				Name: "tags",
				Comment: "Configure tags.",
				Vals: map[string]string{
					"name": "str",
					"memo": "str",
				},
			},
		},
	}
	return info, nil
}

func (p *Provider) List() ([]plug.Tea, error) {
	return p.DB.ListTeas()
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
