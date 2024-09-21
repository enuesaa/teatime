package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TeaSrv struct {
	repos repository.Repos
}

func (srv *TeaSrv) ListTeas(teapod string, teabox string) ([]plug.Tea, error) {
	name := fmt.Sprintf("%s-%s", teapod, teabox)
	list := make([]plug.Tea, 0)

	if err := srv.repos.DB.FindAll(name, bson.D{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

func (srv *TeaSrv) Act(teapod string, name string, vals []plug.Val) (string, error) {
	// provider, err := srv.GetProvider(teapod)
	// if err != nil {
	// 	return "", err
	// }

	// message, err := provider.Act(plug.ActProps{
	// 	Name: name,
	// 	Vals: vals,
	// })
	// if err != nil {
	// 	return "", err
	// }
	// return message, nil
	return "", nil
}
