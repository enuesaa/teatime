package srvtea

import (
	"github.com/enuesaa/teatime/pkg/repository"
)

func New(repos repository.Repos, teapodName string, teaboxName string) Srv {
	srv := Srv{
		repos: repos,
		teapodName: teapodName,
		teaboxName: teaboxName,
	}
	return srv
}

type Srv struct {
	repos repository.Repos
	teapodName string
	teaboxName string
}
