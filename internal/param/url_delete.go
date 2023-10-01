package param

type DeleteUrlRequest struct {
	UserID string `bson:"id"`
	Url    string `bson:"url"`
}

type DeleteUrlResponse struct {
	Message string `bson:"message"`
}
