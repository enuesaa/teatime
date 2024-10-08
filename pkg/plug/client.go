package plug

import (
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewClient(command string) *plugin.Client {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:        "teatime",
		DisableTime: true,
		Level:       hclog.Info,
	})
	config := plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{},
		},
		Logger: logger,
		Cmd:    exec.Command(command),
	}
	return plugin.NewClient(&config)
}

func NewClientProvider(command string) (ProviderInterface, error) {
	client := NewClient(command)

	// TODO: defer client.Kill()
	rpcc, err := client.Client()
	if err != nil {
		return nil, err
	}
	raw, err := rpcc.Dispense("main")
	if err != nil {
		return nil, err
	}
	return raw.(ProviderInterface), nil
}
