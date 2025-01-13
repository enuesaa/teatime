package usecase

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvlog"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func CreateTea(repos repository.Repos, teapodName string, teaboxName string, raw srvtea.Raw) (string, error) {
	teapodSrv := srvteapod.New(repos)
	logSrv := srvlog.New(repos)

	logs, err := teapodSrv.On(teapodName, plug.EventDataCreated, teaboxName, raw)
	// ignore error because this is not critical. also, plugin already executed.
	logSrv.CreateFromPlugLogs(logs)
	if err != nil {
		return "", err
	}

	// TODO
	return "", nil
}
