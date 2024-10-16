package logs

type Item struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}
