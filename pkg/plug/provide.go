package plug

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type ProvideInit = func(db DB, logger Logger) ProviderInterface

func Provide(pinit ProvideInit) {
	teapod := os.Args[0]

	db := DB{}
	db.teapod = teapod

	hclogger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
	logger := Logger{hclogger, teapod}
	provider := pinit(db, logger)

	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				impl:   provider,
				logger: logger,
			},
		},
		Logger: hclogger,
	}
	plugin.Serve(&config)
}
