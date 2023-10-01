package url

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Service) GetUrls(req param.GetUrlRequest) (param.GetUrlResponse, error) {
	var resp param.GetUrlResponse
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			return resp, echo.NewHTTPError(http.StatusNotFound, "user not found")

		}
		return resp, echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}
	resp.Urls = user.Urls
	return resp, nil
}
