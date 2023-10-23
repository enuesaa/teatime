package main

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
)

type Hello struct {}
func (g *Hello) List(args interface{}, resp *string) error {
	*resp = "Hello"
	return nil
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hey",
	MagicCookieValue: "hello",
}

func main() {
	// おそらくデータのスキーマ定義をする
	// プラグインではリクエスト/レスポンスをJSON形式(struct?)で渡すのみ。presenterはteatime本体の責務
	fmt.Println("a")
}

func old() {
	var pluginMap = map[string]plugin.Plugin{
		"hello": &SomethingDataPlugin{},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}