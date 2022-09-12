package model

import (
	"time"
)

type User struct {
	ID            int64     `bson:"_id"`
	Name          string    `bson:"name"`
	Password      string    `bson:"password"`
	Token         string    `bson:"token"`
	Refresh_token string    `bson:"refresh_token"`
	Created_at    time.Time `bson:"created_at"`
	Updated_at    time.Time `bson:"updated_at"`
}
