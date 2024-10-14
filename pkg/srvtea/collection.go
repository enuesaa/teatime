package srvtea

type Tea struct {
	Id string `bson:"_id"`
	Data map[string]interface{} `bson:"data"`
}
