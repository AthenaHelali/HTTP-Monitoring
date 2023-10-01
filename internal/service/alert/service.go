package alert

import "github.com/AthenaHelali/HTTP-Monitoring/internal/model"

type repository interface {
	GetUserByID(id string) (model.User, error)
}

type Service struct {
	repo repository
}

func New(repo repository) Service {
	return Service{repo: repo}
}
