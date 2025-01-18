package plug

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type ProvideInit = func(dbr repository.DBRepository, logger Logger) ProviderInterface

func Provide(pinit ProvideInit) {
	dbr := repository.DBRepository{}
	dbr.Connect()
	defer dbr.Disconnect()

	logger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
	plogger := Logger{logger}
	provider := pinit(dbr, plogger)

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
