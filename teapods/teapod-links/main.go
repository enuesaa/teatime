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

func (p *Provider) List(props plug.ListProps) ([]db.Tea, error) {
	list := []db.Tea{}
	query := p.db.Query(props.Teapod, props.Teabox)

	if err := query.FindAll(bson.M{}, &list, bson.M{}); err != nil {
		return list, err
	}
	return list, nil	
}

func (p *Provider) Get(props plug.GetProps) (db.Tea, error) {
	doc := db.Tea{}
	query := p.db.Query(props.Teapod, props.Teabox)

	id, err := bson.ObjectIDFromHex(props.TeaId)
	if err != nil {
		return doc, err
	}
	filter := bson.M{"_id": id}

	if err := query.Find(filter, &doc); err != nil {
		return doc, err
	}
	p.logger.Info("found: %s", props.TeaId)

	return doc, nil
}

func (p *Provider) Create(props plug.CreateProps) (string, error) {
	if err := ValidateLinkTea(props.Data); err != nil {
		return "", err
	}
	query := p.db.Query(props.Teapod, props.Teabox)

	teaId, err := query.Create(props.Data)
	if err != nil {
		return "", err
	}
	return teaId, nil
}

func (p *Provider) Update(props plug.UpdateProps) (string, error) {
	if err := ValidateLinkTea(props.Data); err != nil {
		return "", err
	}

	query := p.db.Query(props.Teapod, props.Teabox)

	teaId, err := query.Update(props.TeaId, props.Data)
	if err != nil {
		return "", err
	}
	return teaId, nil
}

func (p *Provider) Delete(props plug.DeleteProps) (bool, error) {
	query := p.db.Query(props.Teapod, props.Teabox)

	if err := query.Delete(props.TeaId); err != nil {
		return false, err
	}
	return true, nil
}
