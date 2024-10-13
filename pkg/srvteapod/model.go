package srvteapod

var ModelName = "teapods"

type Teapod struct {
	Name string `bson:"name"`
}
