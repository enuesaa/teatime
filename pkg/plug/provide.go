package plug

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type ProvideInit = func(db DB, logger Logger) ProviderInterface

func Provide(pinit ProvideInit) {
	hclogger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
	logger := Logger{hclogger}
	provider := pinit(DB{}, logger)

	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: provider,
				logger: logger,
			},
		},
		Logger: hclogger,
	}
	plugin.Serve(&config)
}
