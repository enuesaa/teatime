package repository

type Repos struct {
	Log LogRepositoryInterface
	DB  DBRepositoryInterface
}

func New() Repos {
	return Repos{
		Log: &LogRepository{},
		DB: &DBRepository{},
	}
}
