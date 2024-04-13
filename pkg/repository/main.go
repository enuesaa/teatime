package repository

type Repos struct {
	DB  DBRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		DB:  &DBRepository{},
		Log: &LogRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		DB:  &DBRepository{},
		Log: &LogRepository{},
	}
}
