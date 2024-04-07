package plug

import (
	"context"
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

type DB struct {
	teapod string
	repos repository.Repos
}

func (d *DB) Init(teapod string) error {
	d.teapod = teapod
	d.repos = repository.New()
	return d.repos.DB.Open()
}

func (d *DB) Close() error {
	return nil
}
func (d *DB) ListTeas() ([]Tea, error) {
	query, err := d.repos.DB.Query()
	if err != nil {
		return make([]Tea, 0), err
	}
	dbteas, err := query.ListTeas(context.Background(), d.teapod)
	if err != nil {
		return make([]Tea, 0), err
	}

	list := make([]Tea, 0)
	for _, dbtea := range dbteas {
		list = append(list, Tea{
			Teaid: dbtea.Teaid,
		})
	}
	return list, err
}

func (d *DB) GetTea(teaid string) (Tea, error) {
	query, err := d.repos.DB.Query()
	if err != nil {
		return Tea{}, err
	}
	param := dbq.GetTeaParams{
		Teapod: d.teapod,
		Teaid: teaid,
	}
	dbtea, err := query.GetTea(context.Background(), param)
	if err != nil {
		return Tea{}, err
	}
	var vals map[string]string
	if err := json.Unmarshal([]byte(dbtea.Value.(string)), &vals); err != nil {
		return Tea{}, err
	}
	return Tea{Teaid: dbtea.Teaid, Vals: vals}, nil
}

func (d *DB) CreateTea(tea Tea) error {
	query, err := d.repos.DB.Query()
	if err != nil {
		return err
	}
	valuebytes, err := json.Marshal(tea.Vals)
	if err != nil {
		return err
	}
	param := dbq.CreateTeaParams{
		Teapod: d.teapod,
		Collection: "",
		Teaid: tea.Teaid,
		Value: string(valuebytes),
	}
	_, err = query.CreateTea(context.Background(), param)
	return err
}

func (d *DB) DeleteTea(teaid string) error {
	query, err := d.repos.DB.Query()
	if err != nil {
		return err
	}
	param := dbq.DeleteTeaParams{
		Teapod: d.teapod,
		Teaid: teaid,
	}
	return query.DeleteTea(context.Background(), param)
}
