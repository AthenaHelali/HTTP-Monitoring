package url

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (s Service) CreateUrl(req param.CreateUrlRequest) (param.CreateUrlResponse, error) {
	var resp param.CreateUrlResponse

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			return resp, echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return resp, echo.NewHTTPError(http.StatusInternalServerError, "user not found")
	}

	if len(user.Urls) >= 20 {
		return resp, echo.NewHTTPError(http.StatusInternalServerError, errors.New("you can just add 20 urls"))
	}

	for _, x := range user.Urls {
		if x.URL == req.Url {
			return resp, echo.NewHTTPError(http.StatusBadRequest, errors.New("this url already exists"))
		}
	}

	url := model.URL{
		URL:       req.Url,
		Threshold: 20,
		Failed:    0,
		Succeeded: 0,
		CreatedAt: time.Now(),
	}
	user.Urls = append(user.Urls, url)

	s.repo.ReplaceUser(user)
	if err != nil {
		return resp, echo.NewHTTPError(http.StatusInternalServerError, "Error in add url")
	}

	resp.CreatedAt = url.CreatedAt
	resp.Url = url.URL

	return resp, nil
}
