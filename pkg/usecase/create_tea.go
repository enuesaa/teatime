package usecase

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvlog"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func CreateTea(repos repository.Repos, teapodName string, teaboxName string, raw srvtea.Raw) (string, error) {
	teapodSrv := srvteapod.New(repos)
	teaSrv := srvtea.New(repos, teapodName, teaboxName)
	logSrv := srvlog.New(repos)

	meta := map[string]string{
		"teabox": teaboxName,
	}
	logs, err := teapodSrv.On(teapodName, "data.created", meta, raw)
	if err != nil {
		return "", err
	}
	// ignore error because this is not critical and also, plugin already executed.
	logSrv.CreateFromPlugLogs(logs)

	return teaSrv.Create(raw)
}
