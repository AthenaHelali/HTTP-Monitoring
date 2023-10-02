package uservalidator

import "github.com/AthenaHelali/HTTP-Monitoring/internal/model"

const (
	IDRegex = "^[0-9]{8}$"
)

type Repository interface {
	IsIdUnique(id string) (bool, error)
	GetUserByID(id string) (model.User, error)
}
type Validator struct {
	repo Repository
}

func New(repository Repository) Validator {
	return Validator{repo: repository}
}
