package srvtea

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Raw map[string]interface{}

type Tea struct {
	InternalId *bson.ObjectID `bson:"_id"`
	Created    time.Time      `bson:"created"`
	Updated    time.Time      `bson:"updated"`
	Data       Raw            `bson:"data"`
}

func (tea *Tea) Id() string {
	if tea.InternalId == nil {
		return ""
	}
	return tea.InternalId.Hex()
}

type CreateData struct {
	Created time.Time `bson:"created"`
	Updated time.Time `bson:"updated"`
	Data Raw `bson:"data"`
}
type UpdateData struct {
	Updated time.Time `bson:"updated"`
	Data Raw `bson:"data"`
}

func (srv *Srv) CollectionName() string {
	return fmt.Sprintf("%s-%s", srv.teapodName, srv.teaboxName)
}
