package user

import (
	"fmt"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	pass := []byte(req.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	createdUser, err := s.repo.RegisterUser(&model.User{
		ID:        req.ID,
		Name:      req.Name,
		Password:  string(hashedPass),
		CreatedAt: time.Time{},
		Alerts:    []model.History{},
		History:   []model.History{},
		Urls:      []model.URL{},
	})

	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	resp := param.RegisterResponse{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		CreatedAt: createdUser.CreatedAt,
	}

	return resp, nil
}
