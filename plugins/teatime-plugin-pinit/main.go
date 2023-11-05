package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

func main() {
	logger := plug.NewServeLogger()
	logger.Info("bbb")

	connector := plug.Connector{
		Impl: &Handler{},
	}
	plugin.Serve(plug.NewServeConfig(connector))
}
