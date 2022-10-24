package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/request"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/store"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/zap"
)

type App struct {
	Store  store.UserMongodb
	Logger *zap.Logger
}

func (u App) GetAll(c echo.Context) error {
	ss, err := u.Store.GetAll(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, ss)
}

type JWTCustomClaims struct {
	ID string `json:"name"`
	jwt.StandardClaims
}

func (u App) login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		u.Logger.Error("invalid login request", zap.Error(err))

		return echo.ErrBadRequest
	}

	user, err := u.Store.Get(c.Request().Context(), req.Username)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			u.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	if user.Password != req.Password {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JWTCustomClaims{
		req.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (s App) SignUp(c echo.Context) error {
	var req request.User

	if err := c.Bind(&req); err != nil {
		body, _ := io.ReadAll(c.Request().Body)

		s.Logger.Error("can't build request to user",
			zap.Error(err),
			zap.ByteString("body", body),
		)

		return echo.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		s.Logger.Error("request validation faild",
			zap.Error(err),
			zap.Any("request", req),
		)
		return echo.ErrBadRequest
	}
	u := &model.User{
		ID:         req.ID,
		Name:       req.Name,
		Password:   req.Password,
		Created_at: time.Now(),
	}
	if err := s.Store.Save(c.Request().Context(), u); err != nil {
		var errDuplicate store.DuplicateUserError
		if ok := errors.As(err, &errDuplicate); ok {
			s.Logger.Error("duplicate student",
				zap.Error(err),
				zap.String("id", u.ID),
			)
			return echo.ErrBadRequest
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, u)
}

func (u App) GetHistory(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	ctx := c.Request().Context()
	user, err := u.Store.Get(ctx, claims.ID)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			u.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, user.History)
}
func (u App) GetAlerts(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	ctx := c.Request().Context()
	user, err := u.Store.Get(ctx, claims.ID)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			u.Logger.Error("user not found", zap.Error(err))
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, user.Alerts)
}

func (u App) Register(e *echo.Echo) {
	//g.Use(middleware.JWTWithConfig(middleware.JWTConfig{

	//}))
	e.POST("/signup", u.SignUp)
	e.POST("/login", u.login)

	restricted := e.Group("monitoring")
	config := middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	restricted.Use(middleware.JWTWithConfig(config))
	restricted.GET("/history", u.GetHistory)
	restricted.GET("/alerts", u.GetAlerts)
	restricted.GET("/all", u.GetAll)
	restricted.POST("/create-url", u.CreateUrl)
	restricted.GET("/get-urls", u.GetUrls)
	restricted.DELETE("/delete-url", u.DeleteUrl)

}

func (app App) Start() {
	go func() {
		app.MonitorAllUsers()
	}()
}

func (app App) CreateUrl(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := app.Store.Get(c.Request().Context(), claims.ID)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			app.Logger.Error("user not found", zap.Error(err))
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
	app.Store.Replace(c.Request().Context(), user)
	if err != nil {
		app.Logger.Error("Error in add url", zap.Error(err))
		return echo.ErrBadRequest
	}
	app.RequestHTTP(user.ID, url, c.Request().Context())
	return c.JSON(http.StatusOK, user)
}

func (app App) DeleteUrl(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := app.Store.Get(c.Request().Context(), claims.ID)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			app.Logger.Error("user not found", zap.Error(err))
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
	app.Store.Replace(c.Request().Context(), user)
	if err != nil {
		app.Logger.Error("Error in delete url", zap.Error(err))
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, "URL deleted for user successfully")
}

func (u App) GetUrls(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)
	user, err := u.Store.Get(c.Request().Context(), claims.ID)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			u.Logger.Error("user not found", zap.Error(err))
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

func (app App) RequestHTTP(id string, url model.URL, ctx context.Context) error {
	// var user model.User
	user, err := app.Store.Get(ctx, id)
	if err != nil {
		var ErrNotFound store.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			app.Logger.Error("user not found", zap.Error(err))
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
		app.Logger.Error("Error in get response", zap.Error(err))
		resp = &http.Response{StatusCode: 404}
	}
	fmt.Printf("url: %s  statuscode: %d\n", url.URL, resp.StatusCode)
	var history model.History
	history.URL = url
	history.Status_Code = resp.StatusCode
	history.Request_time, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
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
	app.Store.Replace(ctx, user)
	return nil
	// return c.JSON(http.StatusOK, user)
}

func (u App) MonitorAllUsers() {
	for {
		users, err := u.Store.GetAll(context.Background())
		if err == nil {
			for _, user := range users {
				for _, x := range user.Urls {
					fmt.Printf("url: %s  userId: %s\n", x.URL, user.ID)
					u.RequestHTTP(user.ID, x, context.Background())
				}
			}
		}
		time.Sleep(8 * time.Second)
	}
}
