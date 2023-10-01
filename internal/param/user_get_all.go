package param

import "github.com/AthenaHelali/HTTP-Monitoring/internal/model"

type GetUsersRequest struct {
	//TODO add token

}

type GetUsersResponse struct {
	Users []model.User
}
