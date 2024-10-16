package srvlog

import (
	"github.com/enuesaa/teatime/pkg/repository"
)

func New(repos repository.Repos) Srv {
	srv := Srv{
		repos: repos,
	}
	return srv
}

type Srv struct {
	repos repository.Repos
}
