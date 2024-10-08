package repository

type Repos struct {
	Log LogRepositoryInterface
	DB  DBRepositoryInterface
}

func (repos *Repos) Startup() error {
	return repos.DB.Connect()
}

func (repos *Repos) End() error {
	return repos.DB.Disconnect()
}

func New() Repos {
	return Repos{
		Log: &LogRepository{},
		DB:  &DBRepository{},
	}
}
