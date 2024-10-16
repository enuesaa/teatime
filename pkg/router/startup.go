package router

import (
	"time"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvlog"
)

func PutStartupLog(repos repository.Repos) error {
	logSrv := srvlog.New(repos)
	_, err := logSrv.Create(srvlog.LogMessage{
		Created: time.Now().Format(time.RFC3339),
		Message: "teatime app started",
	})

	return err
}
