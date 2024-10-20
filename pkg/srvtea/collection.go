package srvtea

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Raw map[string]interface{}

type Tea struct {
	InternalId *bson.ObjectID `bson:"_id"`
	CreatedAt  time.Time      `bson:"createdAt"`
	UpdatedAt  time.Time      `bson:"updatedAt"`
	Data       Raw            `bson:"data"`
}

func (tea *Tea) Id() string {
	if tea.InternalId == nil {
		return ""
	}
	return tea.InternalId.Hex()
}

type CreateData struct {
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	Data Raw `bson:"data"`
}
type UpdateData struct {
	UpdatedAt time.Time `bson:"updatedAt"`
	Data Raw `bson:"data"`
}

func (srv *Srv) CollectionName() string {
	return fmt.Sprintf("%s-%s", srv.teapodName, srv.teaboxName)
}
