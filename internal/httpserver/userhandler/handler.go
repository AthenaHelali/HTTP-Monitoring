package userhandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/user"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/validator/uservalidator"
)

type Handler struct {
	UserSvc       user.Service
	UserValidator uservalidator.Validator
}
