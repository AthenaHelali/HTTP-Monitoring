package param

import "time"

type CreateUrlRequest struct {
	UserID string `bson:"id"`
	Url    string `bson:"url"`
}

type CreateUrlResponse struct {
	Url       string
	CreatedAt time.Time
}
