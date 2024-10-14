package srvtea

import "go.mongodb.org/mongo-driver/v2/bson"

type Raw map[string]interface{}

func (raw *Raw) ToTea() Tea {
	return Tea{
		Data: *raw,
	}
}

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

func (srv *Srv) CollectionName() string {
	return srv.teaboxName
}
