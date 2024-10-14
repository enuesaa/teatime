package usecase

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvlog"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func CreateTea(repos repository.Repos, teapodName string, teaboxName string, document plug.M) (string, error) {
	teapodSrv := srvteapod.New(repos)
	teaSrv := srvtea.New(repos, teapodName, teaboxName)
	logSrv := srvlog.New(repos)

	logs, err := teapodSrv.On(teapodName, "data.created", []plug.Tea{})
	if err != nil {
		return "", err
	}
	// ignore error because this is not critical and also, plugin already executed.
	logSrv.CreateFromPlugLogs(logs)

	return teaSrv.Create(document)
}
