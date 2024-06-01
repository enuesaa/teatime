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
				ValDefs: []plug.ValDef{
					{Name: "title", Cast: plug.ValCastStr, Nullable: false},
					{Name: "link", Cast: plug.ValCastStr, Nullable: false},
					{Name: "memo", Cast: plug.ValCastStr, Nullable: false},
					{Name: "priority", Cast: plug.ValCastNum, Nullable: false},
				},
			},
			{
				Name: "tags",
				Comment: "Configure tags.",
				ValDefs: []plug.ValDef{
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
	return make([]plug.Tea, 0), nil
}

func (p *Provider) Act(props plug.ActProps) (string, error) {
	return "", nil
}
