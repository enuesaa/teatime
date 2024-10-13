package srvtea

import (
	"github.com/enuesaa/teatime/pkg/repository"
)

func New(repos repository.Repos, teapod string, teabox string) Srv {
	srv := Srv{
		repos:    repos,
		teapod:   teapod,
		teabox:   teabox,
	}
	return srv
}

type Srv struct {
	repos    repository.Repos
	teapod   string
	teabox   string
}
