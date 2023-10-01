package model

import "time"

type URL struct {
	URL       string    `bson:"url"`
	Threshold int       `bson:"threshold"`
	Failed    int       `bson:"failed"`
	Succeeded int       `bson:"succeeded"`
	CreatedAt time.Time `bson:"created_at"`
}
