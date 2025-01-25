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

func (q *TeapodQuery) FindAll(filter M, sort M) ([]Teapod, error) {
	var res []Teapod
	if err := q.query.findAll(filter, &res, sort); err != nil {
		return res, err
	}
	return res, nil
}

func (q *TeapodQuery) Find(filter M) (Teapod, error) {
	var res Teapod
	if err := q.query.find(filter, &res); err != nil {
		return res, err
	}
	return res, nil
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
