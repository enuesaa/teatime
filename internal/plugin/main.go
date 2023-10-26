package plugin

import (
	"fmt"
	"net/rpc"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

type SomethingDataClient struct{ client *rpc.Client }

func (g *SomethingDataClient) List() string {
	var resp string
	g.client.Call("Plugin.List", new(interface{}), &resp)
	return resp
}

type SomethingDataPlugin struct{}

func (p *SomethingDataPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, nil
}
func (SomethingDataPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &SomethingDataClient{client: c}, nil
}

func main() {
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "hey",
		MagicCookieValue: "hello",
	}

	var pluginMap = map[string]plugin.Plugin{
		"hello": &SomethingDataPlugin{}, //
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("../plugin/plugin"),
	})
	defer client.Kill()

	rpcc, _ := client.Client()
	raw, _ := rpcc.Dispense("hello")
	res, _ := raw.(PluginInterface).List()
	fmt.Println(res)
}
