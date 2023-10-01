package user

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
)

func (s Service) GetAll() (param.GetUsersResponse, error) {
	resp := param.GetUsersResponse{}
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return resp, err
	}
	resp.Users = users
	return resp, nil

}
