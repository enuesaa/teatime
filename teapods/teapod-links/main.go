package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func main() {
	provider := Provider{}
	plug.Serve(&provider, "links")
}

type Provider struct {
	*plug.Provider
}

func (p *Provider) Info() plug.InfoResult {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
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
		list = append(list, tea.Rid)
	}
	return p.NewListResult(list)
}

func (p *Provider) Get(rid string) plug.GetResult {
	_, err := p.DBGetTea(rid)
	if err != nil {
		return p.NewGetErr(err)
	}
	row := plug.Tea{
		Rid: rid,
		Value: make(plug.Value, 0),
	}
	return p.NewGetResult(row)
}

func (p *Provider) Set(row plug.Tea) error {
	err := p.DBCreateTea("id", "")
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) Del(rid string) error {
	return nil
}

func (p *Provider) GetCard(name string) plug.GetCardResult {
	return p.NewGetCardErr(nil)
}
