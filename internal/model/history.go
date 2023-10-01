package model

import "time"

type History struct {
	URL         URL       `bson:"url"`
	StatusCode  int       `bson:"status_code"`
	RequestTime time.Time `bson:"request_time"`
}
