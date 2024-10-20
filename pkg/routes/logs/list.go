package logs

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvlog"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)

	logSrv := srvlog.New(cc.Repos)
	logs, err := logSrv.List()
	if err != nil {
		return err
	}

	list := []Item{}
	for _, l := range logs {
		list = append(list, Item{
			Created: l.Created,
			Message: l.Message,
		})
	}

	return cc.WithItems(list)
}
