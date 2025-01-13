package usecase

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvlog"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func UpdateTea(repos repository.Repos, teapodName string, teaboxName string, teaId string, raw srvtea.Raw) (string, error) {
	teapodSrv := srvteapod.New(repos)
	teaSrv := srvtea.New(repos, teapodName, teaboxName)
	logSrv := srvlog.New(repos)

	logs, err := teapodSrv.On(teapodName, plug.EventDataCreated, teaboxName, raw)
	// ignore error because this is not critical and also, plugin already executed.
	logSrv.CreateFromPlugLogs(logs)
	if err != nil {
		return "", err
	}

	return teaSrv.Update(teaId, raw)
}
