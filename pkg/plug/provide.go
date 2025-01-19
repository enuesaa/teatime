package plug

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type ProvideInit = func(db DB, logger Logger) ProviderInterface

func Provide(pinit ProvideInit) {
	db := DB{}
	db.teapod = os.Args[0]

	hclogger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
	logger := Logger{hclogger}
	provider := pinit(db, logger)

	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				impl: provider,
				logger: logger,
			},
		},
		Logger: hclogger,
	}
	plugin.Serve(&config)
}
