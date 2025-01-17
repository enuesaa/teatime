package srvlog

type LogMessage struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}

func (srv *Srv) CollectionName() string {
	return "logs"
}
