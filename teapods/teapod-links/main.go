package main

import (
	"context"
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
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
	kvs, err := p.Query.ListTeas(ctx, "links")
	if err != nil {
		return plug.NewListErrResult(err)
	}
	fmt.Println(kvs)

	list := make([]string, 0)
	return plug.NewListResult(list)
}

func (h *Provider) Get(id string) plug.GetResult {
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
