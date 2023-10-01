package userhandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/user"
)

type Handler struct {
	UserSvc user.Service
}
