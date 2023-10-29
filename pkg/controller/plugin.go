package controller

import (
	"fmt"
	"os/exec"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/gin-gonic/gin"

	"github.com/hashicorp/go-plugin"
)

func InvokePlugin(c *gin.Context) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plug.NewHandshakeConfig(),
		Plugins: map[string]plugin.Plugin{
			"main": &plug.PluginConnector{},
		},
		Cmd: exec.Command("./plugins/teatime-plugin-pinit/teatime-plugin-pinit"),
	})
	defer client.Kill()

	rpcc, err := client.Client()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	raw, err := rpcc.Dispense("main")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	list := raw.(plug.PluginInterface).Resources()
	for _, resource := range list {
		fmt.Printf("%+v\n", resource)
		fmt.Printf("%+v\n", resource.Schema().StringValue)
	}
}
