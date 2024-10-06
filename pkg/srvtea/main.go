package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func New(repos repository.Repos, teapod string, teabox string) (Srv, error) {
	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return Srv{}, err
	}
	srv := Srv{
		repos:    repos,
		teapod:   teapod,
		teabox:   teabox,
		provider: provider,
	}
	return srv, nil
}

type Srv struct {
	repos    repository.Repos
	teapod   string
	teabox   string
	provider plug.ProviderInterface
}
