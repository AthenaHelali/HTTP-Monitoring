package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository/mongo"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/userhandler"
	"net/http"
	"time"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/zap"
)

type App struct {
	Store       mongo.UserMongodb
	Logger      *zap.Logger
	UserHandler userhandler.Handler
}

type JWTCustomClaims struct {
	ID string `json:"name"`
	jwt.StandardClaims
}

func (a App) GetHistory(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := a.Store.GetUserByID(claims.ID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			a.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, user.History)
}
func (a App) GetAlerts(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := a.Store.GetUserByID(claims.ID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			a.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, user.Alerts)
}

func (a App) Register(e *echo.Echo) {
	a.UserHandler.SetUserRoutes(e)

	restricted := e.Group("monitoring")
	config := middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	restricted.Use(middleware.JWTWithConfig(config))
	restricted.GET("/history", a.GetHistory)
	restricted.GET("/alerts", a.GetAlerts)
	restricted.POST("/create-url", a.CreateUrl)
	restricted.GET("/get-urls", a.GetUrls)
	restricted.DELETE("/delete-url", a.DeleteUrl)

}

func (a App) Start() {
	go func() {
		a.MonitorAllUsers()
	}()
}

func (a App) CreateUrl(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := a.Store.GetUserByID(claims.ID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			a.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	var url model.URL
	if err := c.Bind(&url); err != nil {
		return echo.ErrNotFound
	}
	if len(user.Urls) >= 20 {
		c.JSON(http.StatusInternalServerError, errors.New("you can just add 20 urls"))
	}
	var result bool = false
	for _, x := range user.Urls {
		if x.URL == url.URL {
			result = true
			break
		}
	}
	if result {
		c.JSON(http.StatusInternalServerError, errors.New("this url already exists"))
	}
	url.Failed = 0
	user.Urls = append(user.Urls, url)
	a.Store.Replace(c.Request().Context(), user)
	if err != nil {
		a.Logger.Error("Error in add url", zap.Error(err))
		return echo.ErrBadRequest
	}
	a.RequestHTTP(user.ID, url, c.Request().Context())
	return c.JSON(http.StatusOK, user)
}

func (a App) DeleteUrl(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := a.Store.GetUserByID(claims.ID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			a.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	// var url model.URL
	// if err := c.Bind(&url); err != nil {
	// 	return echo.ErrNotFound
	// }
	url := "https://" + c.QueryParam("url")
	var result bool = false
	var index int = 0
	for i, x := range user.Urls {
		if x.URL == url {
			result = true
			index = i
			break
		}
	}
	if !result {
		return echo.ErrBadRequest
	}
	user.Urls = append(user.Urls[:index], user.Urls[index+1:]...)
	a.Store.Replace(c.Request().Context(), user)
	if err != nil {
		a.Logger.Error("Error in delete url", zap.Error(err))
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, "URL deleted for user successfully")
}

func (a App) GetUrls(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := a.Store.GetUserByID(claims.ID)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			a.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	var simpleURLs []string
	for _, url := range user.Urls {
		simpleURLs = append(simpleURLs, url.URL)
	}
	return c.JSON(http.StatusOK, simpleURLs)
}

func (a App) RequestHTTP(id string, url model.URL, ctx context.Context) error {
	// var user model.User
	user, err := a.Store.GetUserByID(id)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			a.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	var isUrlExist bool = false
	var index int = 0
	for i, x := range user.Urls {
		if x.URL == url.URL {
			isUrlExist = true
			index = i
			break
		}
	}
	if !isUrlExist {
		return fmt.Errorf("url doesn't belong to the user")
	}
	resp, err := http.Get(url.URL)
	if err != nil {
		a.Logger.Error("Error in get response", zap.Error(err))
		resp = &http.Response{StatusCode: 404}
	}
	fmt.Printf("url: %s  statuscode: %d\n", url.URL, resp.StatusCode)
	var history model.History
	history.URL = url
	history.StatusCode = resp.StatusCode
	history.RequestTime, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		user.Urls[index].Failed++
		if user.Urls[index].Failed == user.Urls[index].Threshold {
			user.Alerts = append(user.Alerts, history)
			user.Urls[index].Failed = 0
		}
	} else {
		user.Urls[index].Succeeded++
	}
	user.History = append(user.History, history)
	a.Store.Replace(ctx, user)
	return nil
	// return c.JSON(http.StatusOK, user)
}

func (a App) MonitorAllUsers() {
	for {
		users, err := a.Store.GetAllUsers()
		if err == nil {
			for _, user := range users {
				for _, x := range user.Urls {
					fmt.Printf("url: %s  userId: %s\n", x.URL, user.ID)
					a.RequestHTTP(user.ID, x, context.Background())
				}
			}
		}
		time.Sleep(8 * time.Second)
	}
}
