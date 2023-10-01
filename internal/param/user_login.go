package param

type LoginRequest struct {
	ID       string `bson:"id"`
	Password string `bson:"password"`
}

type LoginResponse struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
	//Token Tokens `bson:"token"`
}
