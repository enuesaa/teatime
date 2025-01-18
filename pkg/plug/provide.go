package plug

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type ProvideInit = func(repos repository.Repos, logger Logger) ProviderInterface

func Provide(pinit ProvideInit) {
	repos := repository.New()
	repos.Startup()
	defer repos.End()

	logger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
	plogger := Logger{logger}
	provider := pinit(repos, plogger)

	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: provider,
			},
		},
		Logger: logger,
	}
	plugin.Serve(&config)
}

type Logger struct {
	hclogger hclog.Logger
}

func (l *Logger) Info(message string) {
	l.hclogger.Info(message)
}
