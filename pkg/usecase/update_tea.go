package usecase

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func UpdateTea(repos repository.Repos, teapodName string, teaboxName string, teaId string, raw srvtea.Raw) (string, error) {
	teapodSrv := srvteapod.New(repos)

	err := teapodSrv.On(teapodName, plug.EventDataCreated, teaboxName, raw)
	if err != nil {
		return "", err
	}
	return teaId, nil
}
