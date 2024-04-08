package repository

type Repos struct {
	DB  DbRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		DB:  &DbRepository{},
		Log: &LogRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		DB:  &DbRepository{},
		Log: &LogRepository{},
	}
}
