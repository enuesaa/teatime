package srvmanage

import (
	"github.com/enuesaa/teatime/pkg/repository"
)

func New(repos repository.Repos) Srv {
	return Srv{
		repos: repos,
	}
}

type Srv struct {
	repos repository.Repos
}
