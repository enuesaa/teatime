package srvtea

import "go.mongodb.org/mongo-driver/v2/bson"

type Raw map[string]interface{}

type Tea struct {
	InternalId *bson.ObjectID `bson:"_id"`
	Data Raw `bson:"data"`
}
func (tea *Tea) Id() string {
	if tea.InternalId == nil {
		return ""
	}
	return tea.InternalId.Hex()
}

type Creation struct {
	Data Raw `bson:"data"`
}

func (srv *Srv) CollectionName() string {
	return srv.teaboxName
}
