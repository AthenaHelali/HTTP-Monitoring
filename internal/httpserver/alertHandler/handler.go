package alertHandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/alert"
)

type Handler struct {
	UrlSvc alert.Service
}
