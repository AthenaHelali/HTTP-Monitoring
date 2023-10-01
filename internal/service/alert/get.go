package alert

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Service) GetAlerts(req param.GetAlertRequest) (param.GetAlertResponse, error) {
	var resp param.GetAlertResponse
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			return resp, echo.NewHTTPError(http.StatusNotFound, "user not found")

		}
		return resp, echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}
	resp.Alerts = user.Alerts
	return resp, nil
}
