package model

import "time"

type History struct {
	URL          URL       `bson:"url"`
	Status_Code  int       `bson:"status_code"`
	Request_time time.Time `bson:"request_time"`
}
