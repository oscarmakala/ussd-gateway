package ussd

import (
	"context"
	"ussd-gateway-go/internal/application/core/domain"
)

type Adapter struct {
}

func NewAdapter() (*Adapter, error) {
	return &Adapter{}, nil
}

func (adapter *Adapter) HandleUssdRequest(ctx context.Context, ussdRequest domain.UssdRequest) (domain.UssdResponse, error) {
	return domain.UssdResponse{}, nil
}
