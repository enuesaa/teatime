package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

func main() {
	provider := Provider{}
	plug.Serve(&provider)
}

type Provider struct {
	*plug.Provider
}

func (p *Provider) Info() plug.InfoResult {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
	}
	return plug.NewInfoResult(info)
}

func (p *Provider) List() plug.ListResult {
	ctx := context.Background()
	teas, err := p.Query.ListTeas(ctx, "links")
	if err != nil {
		return plug.NewListErrResult(err)
	}

	list := make([]string, 0)
	for _, tea := range teas {
		list = append(list, strings.TrimPrefix(tea.Resource, "links"))
	}
	return plug.NewListResult(list)
}

func (p *Provider) Get(id string) plug.GetResult {
	ctx := context.Background()
	_, err := p.Query.GetTea(ctx, dbq.GetTeaParams{
		Teapod: "links",
		Resource: fmt.Sprintf("links:%s", id),
	})
	if err != nil {
		return plug.NewGetErrResult(err)
	}
	row := plug.Row{
		Id: id,
		Values: make(plug.Values, 0),
	}
	return plug.NewGetResult(row)
}

func (h *Provider) Set(row plug.Row) error {
	return nil
}

func (h *Provider) Del(id string) error {
	return nil
}

func (h *Provider) GetCard(name string) plug.GetCardResult {
	return plug.NewGetCardErrResult(nil)
}
