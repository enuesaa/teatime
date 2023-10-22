package usecase

import "github.com/enuesaa/teatime-app/backend/repository"

type Appcase struct {
	RedisRepo repository.RedisRepositoryInterface
	RssfeedRepo repository.RssfeedRepositoryInterface
}

func NewAppcase() *Appcase {
	return &Appcase{
		RedisRepo: &repository.RedisRepository{},
		RssfeedRepo: &repository.RssfeedRepository{},
	}
}

func (app *Appcase) ListBoard() string {
	return "board:"
}
func (app *Appcase) GetBoard() string {
	return "board:"
}
func (app *Appcase) CreateBoard() string {
	return "board:"
}
func (app *Appcase) UpdateBoard() string {
	return "board:"
}
func (app *Appcase) DeleteBoard() string {
	return "board:"
}

