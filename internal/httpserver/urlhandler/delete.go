package urlhandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/auth"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (h Handler) DeleteUrl(c echo.Context) error {
	var req param.DeleteUrlRequest
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.Claims)

	req.UserID = claims.UserID

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "cant bind request")
	}

	_, err := h.UrlSvc.DeleteUrl(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cant delete url")
	}
	return echo.NewHTTPError(http.StatusOK, "URL deleted for user successfully")

}
