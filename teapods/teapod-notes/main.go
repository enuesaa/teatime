package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

func main() {
	plug.Provide(NewProvider)
}

func NewProvider(db plug.DB, logger plug.Logger) plug.ProviderInterface {
	return &Provider{db, logger}
}

type Provider struct {
	db plug.DB
	logger plug.Logger
}

func (p *Provider) OnStartup() error {
	return p.db.Connect()
}

func (p *Provider) OnShutdown() error {
	return p.db.Close()
}

func (p *Provider) Info() (plug.Info, error) {
	info := plug.Info{
		Name: "teapod-notes",
		Description: "notes teapod",
		Teaboxes: []plug.Teabox{
			{
				Name: "notes",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "content", Type: plug.TeaboxInputText},
				},
			},
		},
		Actions: []plug.Action{},
	}
	return info, nil
}

func (p *Provider) List(ar plug.ListArgs) ([]db.Tea, error) {
	query := p.db.Use(ar.Teabox)

	return query.FindAll(db.M{}, db.M{})
}

func (p *Provider) Get(ar plug.GetArgs) (db.Tea, error) {
	query := p.db.Use(ar.Teabox)

	return query.Get(ar.TeaId)
}

func (p *Provider) Create(ar plug.CreateArgs) (string, error) {
	if err := ValidateNoteTea(ar.Data); err != nil {
		return "", err
	}
	return p.db.Use(ar.Teabox).Create(ar.Data)
}

func (p *Provider) Update(ar plug.UpdateArgs) (string, error) {
	if err := ValidateNoteTea(ar.Data); err != nil {
		return "", err
	}
	return p.db.Use(ar.Teabox).Update(ar.TeaId, ar.Data)
}

func (p *Provider) Delete(ar plug.DeleteArgs) error {
	return p.db.Use(ar.Teabox).Delete(ar.TeaId)
}

func (p *Provider) Act(ar plug.ActArgs) (string, error) {
	return "", nil
}
