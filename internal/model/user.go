package model

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	Alerts    []Alert   `bson:"alert"`
	History   []History `bson:"history"`
	Urls      []URL     `bson:"urls"`
}
