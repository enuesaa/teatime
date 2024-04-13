package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

var db plug.DB

func main() {
	db = plug.NewDB("links")
	defer db.Close()

	plug.Serve(&Provider{})
}

type Provider struct {}

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

func (p *Provider) List(props plug.ListProps) ([]plug.Tea, error) {
	if props.TeaboxName != nil {
		return db.ListTeasByTeaboxName(*props.TeaboxName)
	}
	return db.ListTeas()
}

func (p *Provider) Get(teaid string) (plug.Tea, error) {
	return db.GetTea(teaid)
}

func (p *Provider) Set(tea plug.Tea) error {
	return db.CreateTea(tea)
}

func (p *Provider) Del(teaid string) error {
	return db.DeleteTea(teaid)
}

func (p *Provider) GetCard(name string) (plug.Card, error) {
	return plug.Card{}, nil
}
