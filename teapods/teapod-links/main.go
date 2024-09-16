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
		Schema: []plug.Valdef{
			{Name: "title", Cast: plug.ValCastStr, Nullable: false},
			{Name: "link", Cast: plug.ValCastStr, Nullable: false},
			{Name: "memo", Cast: plug.ValCastStr, Nullable: false},
			{Name: "priority", Cast: plug.ValCastNum, Nullable: false},
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

func (p *Provider) Act(props plug.ActProps) (string, error) {
	return "", nil
}
