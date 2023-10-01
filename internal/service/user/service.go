package user

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
)

type repository interface {
	RegisterUser(m *model.User) (model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
}
type AuthGenerator interface {
	CreateAccessToken(user model.User) (string, error)
	CreateRefreshToken(user model.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo repository
}

func New(authGenerator AuthGenerator, repo repository) Service {
	return Service{auth: authGenerator, repo: repo}
}
