package main

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var name = "links"

func main() {
	plug.Provide(NewProvider)
}

func NewProvider(dbr repository.DBRepository, logger plug.Logger) plug.ProviderInterface {
	return &Provider{dbr, logger}
}

type Provider struct {
	DB repository.DBRepository
	logger plug.Logger
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
	if err := p.DB.Teas(name, props.Teabox).FindAll(bson.M{}, &list, bson.M{}); err != nil {
		return list, err
	}
	return list, nil	
}

func (p *Provider) Get(props plug.GetProps) (db.Tea, error) {
	doc := db.Tea{}
	filter := bson.M{"_id": props.TeaId}
	if err := p.DB.Teas(name, props.Teabox).Find(filter, &doc); err != nil {
		return doc, err
	}
	return doc, nil
}

func (p *Provider) Create(props plug.CreateProps) (string, error) {
	if err := ValidateLinkTea(props.Data); err != nil {
		p.logger.Info(fmt.Sprintf("tea invalid: %v", err.Error()))
		return "", err
	}
	teaId, err := p.DB.Teas(name, props.Teabox).Create(props.Data)
	if err != nil {
		return "", err
	}
	return teaId, nil
}

func (p *Provider) Update(props plug.UpdateProps) (string, error) {
	if err := ValidateLinkTea(props.Data); err != nil {
		p.logger.Info(fmt.Sprintf("tea invalid: %v", err.Error()))
		return "", err
	}
	teaId, err := p.DB.Teas(name, props.Teabox).Update(props.TeaId, props.Data)
	if err != nil {
		return "", err
	}
	return teaId, nil
}

func (p *Provider) Delete(props plug.DeleteProps) (bool, error) {
	if err := p.DB.Teas(name, props.Teabox).Delete(props.TeaId); err != nil {
		return false, err
	}
	return true, nil
}
