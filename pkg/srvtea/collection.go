package srvtea

type Raw map[string]interface{}

func (raw *Raw) ToTea() Tea {
	return Tea{
		Id: "",
		Data: *raw,
	}
}

type Tea struct {
	Id string `bson:"_id"`
	Data Raw `bson:"data"`
}

func (srv *Srv) CollectionName() string {
	return srv.teaboxName
}
