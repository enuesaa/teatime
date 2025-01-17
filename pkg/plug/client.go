package plug

import (
	"encoding/json"
	"io"
	"os/exec"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type LogMessage struct {
	Created string `bson:"created" json:"timestamp"`
	Message string `bson:"message" json:"@message"`
}

type LogWriter struct {
	io.Writer
	repos repository.Repos
}
func (w *LogWriter) Write(p []byte) (int, error) {

	var message LogMessage
	if err := json.Unmarshal(p, &message); err != nil {
		return 0, err
	}
	if _, err := w.repos.DB.Create("logs", message); err != nil {
		return 0, err
	}
	return len(p), nil
}

func NewClient(command string, repos repository.Repos) *plugin.Client {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:        "teatime",
		DisableTime: true,
		Level:       hclog.Info,
		Output:      &LogWriter{repos: repos},
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
	return plugin.NewClient(&config)
}

func NewClientProvider(command string, repos repository.Repos) (ProviderInterface, error) {
	client := NewClient(command, repos)

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
