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

func (q *LogQuery) FindAll(filter M, sort M) ([]Log, error) {
	var res []Log
	if err := q.query.findAll(filter, &res, sort); err != nil {
		return res, err
	}
	return res, nil
}

func (q *LogQuery) Find(filter M) (Log, error) {
	var res Log
	if err := q.query.find(filter, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (q *LogQuery) Create(doc Log) (string, error) {
	return q.query.create(doc)
}

func (q *LogQuery) Update(id string, doc Log) (string, error) {
	return q.query.update(id, doc)
}

func (q *LogQuery) Delete(id string) error {
	return q.query.delete(id)
}

func (q *LogQuery) DeleteMany(filter M) error {
	return q.query.deleteMany(filter)
}
