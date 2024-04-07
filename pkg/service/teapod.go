package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/google/uuid"
)

func NewTeapodSrv(name string) *TeapodSrv {
	return &TeapodSrv{
		Name: name,
	}
}

type TeapodSrv struct {
	Name string
}

func (srv *TeapodSrv) GetProvider() (plug.ProviderInterface, error) {
	command := fmt.Sprintf("teapod-%s", srv.Name)
	return plug.Run(command)
}

func (srv *TeapodSrv) GetInfo() (plug.Info, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Info{}, err
	}
	return provider.Info()
}

func (srv *TeapodSrv) GetTeabox(name string) (plug.Teabox, error) {
	info, err := srv.GetInfo()
	if err != nil {
		return plug.Teabox{}, fmt.Errorf("teabox not found.")
	}
	for _, teabox := range info.Teaboxes {
		if name == teabox.Name {
			return teabox, nil
		}
	}
	return plug.Teabox{}, fmt.Errorf("teabox not found.")
}

func (srv *TeapodSrv) ListTeas(teaboxName string) ([]plug.Tea, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]plug.Tea, 0), err
	}
	props := plug.ListProps{}
	if teaboxName != "" {
		props.TeaboxName = &teaboxName
	}
	return provider.List(props)
}

func (srv *TeapodSrv) GetTea(teaid string) (plug.Tea, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Tea{}, err
	}
	return provider.Get(teaid)
}

func (srv *TeapodSrv) CreateTea(teaboxName string, vals plug.Vals) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return "", err
	}
	teaid := uuid.NewString()
	tea := plug.Tea{
		Teaid: teaid,
		Teabox: teaboxName,
		Vals: vals,
	}
	if err := provider.Set(tea); err != nil {
		return "", err
	}
	return teaid, nil
}

func (srv *TeapodSrv) UpdateTea(teaid string, vals plug.Vals) (string, error) {
	if err := srv.DeleteTea(teaid); err != nil {
		return "", err
	}
	provider, err := srv.GetProvider()
	if err != nil {
		return teaid, err
	}
	tea := plug.Tea{
		Teaid: teaid,
		Vals: vals,
	}
	if err := provider.Set(tea); err != nil {
		return "", err
	}
	return teaid, nil
}

func (srv *TeapodSrv) DeleteTea(teaid string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Del(teaid)
}

func (srv *TeapodSrv) GetCard(name string) (plug.Card, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Card{}, err
	}
	return provider.GetCard(name)
}

func (srv *TeapodSrv) ValidateTeaboxVals(teabox plug.Teabox, vals plug.Vals) error {
	for key := range teabox.Vals {
		if _, ok := vals[key]; !ok {
			return fmt.Errorf("key `%s` is required.", key)
		}
	}

	for key := range vals {
		if _, ok := teabox.Vals[key]; !ok {
			return fmt.Errorf("additional key `%s` is not allowd.", key)
		}
	}
	return nil
}
