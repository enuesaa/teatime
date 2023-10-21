package main

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func main() {
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "hey",
		MagicCookieValue: "hello",
	}

	var pluginMap = map[string]plugin.Plugin{
		"hello": &SomethingDataPlugin{},
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("../plugin/plugin"),
	})
	defer client.Kill()

	rpcc, _ := client.Client()
	raw, _ := rpcc.Dispense("hello")
	res := raw.(SomethingDataPluginInterface).List()
	fmt.Println(res)
}
