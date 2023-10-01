package param

import "time"

type RegisterRequest struct {
	ID       string `bson:"id"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

type RegisterResponse struct {
	ID        string    `bson:"id"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
}
