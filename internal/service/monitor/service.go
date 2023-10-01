package monitor

import "github.com/AthenaHelali/HTTP-Monitoring/internal/model"

type repository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id string) (model.User, error)
	ReplaceUser(user model.User) error
}

type Service struct {
	repo repository
}

func New(repo repository) Service {
	return Service{repo: repo}
}
