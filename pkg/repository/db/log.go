package db

type Log struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}

type LogQuery struct {
	query Query
}

func (q *LogQuery) CreateCollection() error {
	return q.query.createCollection()
}

func (q *LogQuery) DropCollection() error {
	return q.query.dropCollection()
}

func (q *LogQuery) FindAll(filter M, docs *[]Log, sort M) error {
	return q.query.findAll(filter, docs, sort)
}

func (q *LogQuery) Find(filter M, doc *Log) error {
	return q.query.find(filter, doc)
}

func (q *LogQuery) Create(doc Log) (string, error) {
	return q.query.create(doc)
}

func (q *LogQuery) Update(id string, doc Log) (string, error) {
	return q.query.update(id, doc)
}

func (q *LogQuery) Delete(filter M) error {
	return q.query.delete(filter)
}

func (q *LogQuery) DeleteMany(filter M) error {
	return q.query.deleteMany(filter)
}
