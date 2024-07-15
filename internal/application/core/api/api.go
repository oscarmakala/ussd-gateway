package api

import (
	"context"
	"ussd-gateway-go/internal/application/core/domain"
	"ussd-gateway-go/internal/ports"
)

type Application struct {
	db   ports.DBPort
	ussd ports.UssdPort
}

func NewApplication() *Application {
	return &Application{}
}

func (app Application) HandleUssdRequest(ctx context.Context, ussdRequest domain.UssdRequest) {

}
