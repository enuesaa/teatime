package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

func main() {
	connector := plug.Connector{
		Impl: NewHander(),
	}
	plugin.Serve(plug.NewServeConfig(connector))
}
