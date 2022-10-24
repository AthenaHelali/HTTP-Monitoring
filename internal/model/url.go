package model

type URL struct {
	URL       string `bson:"url"`
	Threshold int    `bson:"threshold"`
	Failed    int    `bson:"failed"`
	Succeeded int    `bson:"succeeded"`
}
