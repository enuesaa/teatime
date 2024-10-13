package srvteapod

type Teapod struct {
	Name string `bson:"name"`
}

func (srv *Srv) ModelName() string {
	return "teapods"
}
