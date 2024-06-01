package repository

type Repos struct {
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Log: &LogRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		Log: &LogRepository{},
	}
}
