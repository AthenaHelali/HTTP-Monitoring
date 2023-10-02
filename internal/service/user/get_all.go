package user

import (
	"fmt"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
)

func (s Service) GetAll(req param.GetUsersRequest) (param.GetUsersResponse, error) {
	_, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return param.GetUsersResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	resp := param.GetUsersResponse{}
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return resp, err
	}

	for _, user := range users {
		resp.Users = append(resp.Users, param.UserInfo{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			Urls:      user.Urls,
		})
	}

	return resp, nil
}
