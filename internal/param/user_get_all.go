package param

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"time"
)

type GetUsersRequest struct {
	UserID string `bson:"_id"`
}

type UserInfo struct {
	ID        string      `bson:"_id"`
	Name      string      `bson:"name"`
	CreatedAt time.Time   `bson:"created_at"`
	Urls      []model.URL `bson:"urls"`
}

type GetUsersResponse struct {
	Users []UserInfo
}
