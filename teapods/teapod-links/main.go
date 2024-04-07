package main

import (
	"encoding/json"
	"log"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func main() {
	provider := Provider{}
	if err := provider.Serve(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

type Provider struct {
	plug.Provider
}

func (p *Provider) Serve() error {
	return p.Provider.Serve("links", p, repository.New())
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
	dbteas, err := p.Repos.DB.ListTeas("links")
	if err != nil {
		return p.NewListErr(err)
	}
	list := make([]string, 0)
	for _, dbtea := range dbteas {
		list = append(list, dbtea.Teaid)
	}
	return p.NewListResult(list)
}

func (p *Provider) Get(teaid string) plug.GetResult {
	dbtea, err := p.Repos.DB.GetTea("links", teaid)
	if err != nil {
		return p.NewGetErr(err)
	}
	var value plug.Value
	if err := json.Unmarshal([]byte(dbtea.Value.(string)), &value); err != nil {
		return p.NewGetErr(err)
	}
	return p.NewGetResult(plug.Tea{Teaid: dbtea.Teaid, Value: value})
}

func (p *Provider) Set(tea plug.Tea) plug.SetResult {
	valuebytes, err := json.Marshal(tea.Value)
	if err != nil {
		return p.NewSetErr(err)
	}
	if err := p.Repos.DB.CreateTea("links", tea.Teaid, string(valuebytes)); err != nil {
		return p.NewSetErr(err)
	}
	return p.NewSetResult()
}

func (p *Provider) Del(teaid string) plug.DelResult {
	if err := p.Repos.DB.DeleteTea("links", teaid); err != nil {
		return p.NewDelErr(err)
	}
	return p.NewDelResult()
}

func (p *Provider) GetCard(name string) plug.GetCardResult {
	return p.NewGetCardErr(nil)
}
