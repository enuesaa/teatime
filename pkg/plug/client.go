package plug

import (
	"os/exec"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewClient(command string, repos repository.Repos) (ProviderInterface, error) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:        "teatime",
		DisableTime: true,
		Level:       hclog.Info,
		Output:      &ClientLogWriter{repos: repos},
		JSONFormat:  true,
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
	client := plugin.NewClient(&config)

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
