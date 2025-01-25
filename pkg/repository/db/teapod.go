package db

type Teapod struct {
	Name        string         `bson:"name"`
	Description string         `bson:"description"`
	Teaboxes    []TeapodTeabox `bson:"teaboxes"`
	Actions     []TeapodAction `bson:"actions"`
}
type TeapodTeabox struct {
	Name   string              `bson:"name"`
	Inputs []TeapodTeaboxInput `bson:"inputs"`
}
type TeapodTeaboxInput struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}
type TeapodAction struct {
	Event   string `bson:"event"`
	Comment string `bson:"comment"`
}

type TeapodQuery struct {
	query Query
}

func (q *TeapodQuery) CreateCollection() error {
	return q.query.createCollection()
}

func (q *TeapodQuery) DropCollection() error {
	return q.query.dropCollection()
}

func (q *TeapodQuery) FindAll(filter M, res *[]Teapod, sort M) error {
	return q.query.findAll(filter, res, sort)
}

func (q *TeapodQuery) Find(filter M, res *Teapod) error {
	return q.query.find(filter, res)
}

func (q *TeapodQuery) Get(id string, res *Tea) error {
	return q.query.get(id, res)
}

func (q *TeapodQuery) Create(doc Teapod) (string, error) {
	return q.query.create(doc)
}

func (q *TeapodQuery) Update(id string, doc Teapod) (string, error) {
	return q.query.update(id, doc)
}

func (q *TeapodQuery) Delete(id string) error {
	return q.query.delete(id)
}

func (q *TeapodQuery) DeleteMany(filter M) error {
	return q.query.deleteMany(filter)
}
