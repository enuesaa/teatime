package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func main() {
	provider := Provider{}
	if err := provider.Serve("links", repository.New()); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type Provider struct {
	plug.Provider
}

func (p *Provider) Info() plug.InfoResult {
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
	return p.NewInfoResult(info)
}

func (p *Provider) List() plug.ListResult {
	teas, err := p.DBListTeas()
	if err != nil {
		return p.NewListErr(err)
	}

	list := make([]string, 0)
	for _, tea := range teas {
		list = append(list, tea.Teaid)
	}
	return p.NewListResult(list)
}

func (p *Provider) Get(teaid string) plug.GetResult {
	tea, err := p.DBGetTea(teaid)
	if err != nil {
		return p.NewGetErr(err)
	}
	return p.NewGetResult(tea)
}

func (p *Provider) Set(tea plug.Tea) plug.SetResult {
	if err := p.DBCreateTea(tea.Teaid, tea.Value); err != nil {
		return p.NewSetErr(err)
	}
	return p.NewSetResult()
}

func (p *Provider) Del(teaid string) plug.DelResult {
	if err := p.DBDeleteTea(teaid); err != nil {
		return p.NewDelErr(err)
	}
	return p.NewDelResult()
}

func (p *Provider) GetCard(name string) plug.GetCardResult {
	return p.NewGetCardErr(nil)
}
