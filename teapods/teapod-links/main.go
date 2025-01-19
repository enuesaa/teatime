package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
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
		Name: "teapod-links",
		Description: "links teapod",
		Teaboxes: []plug.Teabox{
			{
				Name: "links",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "link", Type: plug.TeaboxInputText},
				},
			},
			{
				Name: "notes",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "memo", Type: plug.TeaboxInputText},
				},
			},
		},
		Actions: []plug.Action{
			{
				Name:   "action.sync",
				Comment: "sync links",
			},
		},
	}
	return info, nil
}

func (p *Provider) List(ar plug.ListArgs) ([]db.Tea, error) {
	list := []db.Tea{}
	query := p.db.Use(ar.Teabox)

	if err := query.FindAll(bson.M{}, &list, bson.M{}); err != nil {
		return list, err
	}
	return list, nil	
}

func (p *Provider) Get(ar plug.GetArgs) (db.Tea, error) {
	doc := db.Tea{}
	query := p.db.Use(ar.Teabox)

	id, err := bson.ObjectIDFromHex(ar.TeaId)
	if err != nil {
		return doc, err
	}
	filter := bson.M{"_id": id}

	if err := query.Find(filter, &doc); err != nil {
		return doc, err
	}
	return doc, nil
}

func (p *Provider) Create(ar plug.CreateArgs) (string, error) {
	if err := ValidateLinkTea(ar.Data); err != nil {
		return "", err
	}
	return p.db.Use(ar.Teabox).Create(ar.Data)
}

func (p *Provider) Update(ar plug.UpdateArgs) (string, error) {
	if err := ValidateLinkTea(ar.Data); err != nil {
		return "", err
	}
	return p.db.Use(ar.Teabox).Update(ar.TeaId, ar.Data)
}

func (p *Provider) Delete(ar plug.DeleteArgs) error {
	return p.db.Use(ar.Teabox).Delete(ar.TeaId)
}
