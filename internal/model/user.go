package model

import (
	"time"
)

type User struct {
	ID         string    `bson:"_id"`
	Name       string    `bson:"name"`
	Password   string    `bson:"password"`
	Created_at time.Time `bson:"created_at"`
	Alerts     []History `bson:"alert"`
	History    []History `bson:"history"`
	Urls       []URL     `bson:"urls"`
}
