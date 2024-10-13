package srvteapod

import (
	"fmt"
)

func (srv *Srv) Register(teapodName string) error {
	info, err := srv.Info(teapodName)
	if err != nil {
		return fmt.Errorf("failed to fetch teapod %s", teapodName)
	}

	teapod := Teapod {
		Name: teapodName,
	}
	if _, err := srv.repos.DB.Create(ModelName, teapod); err != nil {
		return err
	}

	for _, teabox := range info.Teaboxes {
		if err := srv.repos.DB.CreateCollection(teabox.Name, teabox.Schema.Bson()); err != nil {
			return err
		}
	}

	// TODO: rollback

	return nil
}
