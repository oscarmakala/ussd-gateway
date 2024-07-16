package ports

import (
	"context"
	"ussd-gateway-go/internal/application/core/domain"
)

type UssdPort interface {
	HandleUssdRequest(ctx context.Context, ussdRequest domain.UssdRequest) (domain.UssdResponse, error)
}

