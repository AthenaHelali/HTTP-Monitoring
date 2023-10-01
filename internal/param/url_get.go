package param

import "github.com/AthenaHelali/HTTP-Monitoring/internal/model"

type GetUrlRequest struct {
	UserID string `bson:"id"`
}

type GetUrlResponse struct {
	Urls []model.URL
}
