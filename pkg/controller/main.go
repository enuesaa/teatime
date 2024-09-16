package controller

import "github.com/enuesaa/teatime/pkg/repository"

func New(repos repository.Repos) Ctl {
	return Ctl{
		repos: repos,
	}
}

type Ctl struct {
	repos repository.Repos
}
