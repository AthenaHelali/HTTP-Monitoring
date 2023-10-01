package param

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
)

type GetAlertRequest struct {
	UserID string `bson:"id"`
}

type GetAlertResponse struct {
	Alerts []model.Alert
}
