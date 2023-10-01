package monitor

import (
	"fmt"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"net/http"
	"time"
)

func (s Service) MonitorAllUsers() {
	for {
		users, err := s.repo.GetAllUsers()
		if err == nil {
			for _, user := range users {
				for index, url := range user.Urls {
					fmt.Printf("url: %s  userId: %s\n", url.URL, user.ID)
					s.RequestHTTP(user, url, index)
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func (s Service) RequestHTTP(user model.User, url model.URL, index int) error {
	resp, err := http.Get(url.URL)
	if err != nil {
		resp = &http.Response{StatusCode: 404}
	}
	fmt.Printf("url: %s  statuscode: %d\n", url.URL, resp.StatusCode)
	var history model.History
	var alert model.Alert
	history.URL = url
	history.StatusCode = resp.StatusCode
	history.RequestTime, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		user.Urls[index].Failed++
		if user.Urls[index].Failed == user.Urls[index].Threshold {
			alert.AlertTime = time.Now()
			alert.URL = url
			alert.Failed = 20
			alert.Succeeded = user.Urls[index].Succeeded
			user.Alerts = append(user.Alerts, alert)
			user.Urls[index].Failed = 0
		}
	} else {
		user.Urls[index].Succeeded++
	}
	user.History = append(user.History, history)

	err = s.repo.ReplaceUser(user)
	if err != nil {
		return err
	}
	return nil
}
