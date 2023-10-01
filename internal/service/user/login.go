package user

import (
	"fmt"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	user, err := s.repo.GetUserByID(req.ID)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	if hErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); hErr != nil {
		return param.LoginResponse{}, fmt.Errorf("password is not correct")
	}
	resp := param.LoginResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	return resp, nil
}
