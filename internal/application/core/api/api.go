package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"ussd-gateway-go/internal/application/core/domain"
	"ussd-gateway-go/internal/ports"
)

type Application struct {
	ussd ports.UssdPort
}

func NewApplication(ussd ports.UssdPort) *Application {
	return &Application{
		ussd: ussd,
	}
}

func (a Application) ProcessRequest(ctx context.Context, ussdRequest domain.UssdRequest) (domain.UssdResponse, error) {
	log.Debug("ProcessRequest called")
	return domain.UssdResponse{}, nil
}
