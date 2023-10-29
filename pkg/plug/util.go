package plug

import (
	"github.com/hashicorp/go-plugin"
)

func NewHandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "hey",
		MagicCookieValue: "hello",
	}
}
