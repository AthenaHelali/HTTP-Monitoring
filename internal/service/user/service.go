package user

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
)

type repository interface {
	RegisterUser(m model.User) (model.User, error)
	GetUserByID(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)
}

type Service struct {
	repo repository
}

func New(repo repository) Service {
	return Service{repo: repo}
}
