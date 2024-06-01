package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
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
				Comment: "Resgister your favorite site and look up later.",
				Valdefs: []plug.Valdef{
					{Name: "title", Cast: plug.ValCastStr, Nullable: false},
					{Name: "link", Cast: plug.ValCastStr, Nullable: false},
					{Name: "memo", Cast: plug.ValCastStr, Nullable: false},
					{Name: "priority", Cast: plug.ValCastNum, Nullable: false},
				},
			},
			{
				Name: "tags",
				Comment: "Configure tags.",
				Valdefs: []plug.Valdef{
					{Name: "name", Cast: plug.ValCastStr, Nullable: false},
					{Name: "memo", Cast: plug.ValCastStr, Nullable: false},
				},
			},
		},
		Actions: []plug.Action{
			{
				Name: "remove",
				Comment: "remove tea",
				Teabox: plug.String("tags"),
			},
		},
	}
	return info, nil
}

func (p *Provider) List(props plug.ListProps) ([]plug.Tea, error) {
	list := make([]plug.Tea, 0)

	switch *props.TeaboxName {
	case "links":
		list = append(list, plug.Tea{
			Teaid: "a",
			Teabox: "links",
			Vals: []plug.Val{
				{Name: "title", Value: "a"},
				{Name: "link", Value: "https://example.com"},
				{Name: "memo", Value: "memo"},
				{Name: "priority", Value: "0"},
			},
		})
		return list, nil
	case "tags":
		return list, nil
	}

	return list, nil
}

func (p *Provider) Act(props plug.ActProps) (string, error) {
	return "", nil
}
