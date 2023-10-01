package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)

	userGroup.POST("/login", h.userLogin)

	userGroup.GET("/all", h.allUsers)

	//userGroup.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))

}
