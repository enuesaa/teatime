package plug

import (
	"context"
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

func NewDB(teapod string) DB {
	db := DB {
		teapod: teapod,
		repos: repository.New(),
	}
	return db
}

type DB struct {
	teapod string
	repos  repository.Repos
}

func (d *DB) openAnyway() error {
	if !d.repos.DB.IsOpen() {
		return d.repos.DB.Open()
	}
	return nil
}

func (d *DB) Close() error {
	return nil
}
func (d *DB) ListTeas() ([]Tea, error) {
	if err := d.openAnyway(); err != nil {
		return make([]Tea, 0), err
	}
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
		var vals map[string]string
		if err := json.Unmarshal([]byte(dbtea.Value.(string)), &vals); err != nil {
			return make([]Tea, 0), err
		}
		list = append(list, Tea{
			Teaid:  dbtea.Teaid,
			Teabox: dbtea.Teabox,
			Vals:   vals,
		})
	}
	return list, err
}

func (d *DB) ListTeasByTeaboxName(teaboxName string) ([]Tea, error) {
	if err := d.openAnyway(); err != nil {
		return make([]Tea, 0), err
	}
	query, err := d.repos.DB.Query()
	if err != nil {
		return make([]Tea, 0), err
	}
	arg := dbq.ListTeasByTeaboxNameParams{
		Teapod: d.teapod,
		Teabox: teaboxName,
	}
	dbteas, err := query.ListTeasByTeaboxName(context.Background(), arg)
	if err != nil {
		return make([]Tea, 0), err
	}

	list := make([]Tea, 0)
	for _, dbtea := range dbteas {
		var vals map[string]string
		if err := json.Unmarshal([]byte(dbtea.Value.(string)), &vals); err != nil {
			return make([]Tea, 0), err
		}
		list = append(list, Tea{
			Teaid:  dbtea.Teaid,
			Teabox: dbtea.Teabox,
			Vals:   vals,
		})
	}
	return list, err
}

func (d *DB) GetTea(teaid string) (Tea, error) {
	if err := d.openAnyway(); err != nil {
		return Tea{}, err
	}
	query, err := d.repos.DB.Query()
	if err != nil {
		return Tea{}, err
	}
	param := dbq.GetTeaParams{
		Teapod: d.teapod,
		Teaid:  teaid,
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
	if err := d.openAnyway(); err != nil {
		return err
	}
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
		Teabox: tea.Teabox,
		Teaid:  tea.Teaid,
		Value:  string(valuebytes),
	}
	_, err = query.CreateTea(context.Background(), param)
	return err
}

func (d *DB) DeleteTea(teaid string) error {
	if err := d.openAnyway(); err != nil {
		return err
	}
	query, err := d.repos.DB.Query()
	if err != nil {
		return err
	}
	param := dbq.DeleteTeaParams{
		Teapod: d.teapod,
		Teaid:  teaid,
	}
	return query.DeleteTea(context.Background(), param)
}
