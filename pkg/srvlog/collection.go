package srvlog

type Log struct {
	Created string
	Message string
}

func (srv *Srv) CollectionName() string {
	return "logs"
}
