package usecase

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func CreateTea(repos repository.Repos, teapodName string, teaboxName string, document plug.M) (string, error) {
	teapodSrv := srvteapod.New(repos)
	if err := teapodSrv.On(teapodName, "data.created", []plug.Tea{}); err != nil {
		return "", err
	}
	teaSrv := srvtea.New(repos, teapodName, teaboxName)
	
	return teaSrv.Create(document)
}
