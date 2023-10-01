package model

import "time"

type Alert struct {
	URL       URL       `bson:"url"`
	Failed    int       `bson:"failed"`
	Succeeded int       `bson:"succeeded"`
	AlertTime time.Time `bson:"alert_time"`
}
