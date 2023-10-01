package url

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Service) DeleteUrl(req param.DeleteUrlRequest) (param.DeleteUrlResponse, error) {
	var resp param.DeleteUrlResponse
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			return resp, echo.ErrNotFound
		}
		return resp, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var result bool = false
	var index int = 0
	for i, x := range user.Urls {
		if x.URL == req.Url {
			result = true
			index = i
			break
		}
	}
	if !result {
		return resp, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Urls = append(user.Urls[:index], user.Urls[index+1:]...)
	s.repo.ReplaceUser(user)
	if err != nil {
		return resp, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return resp, nil
}
