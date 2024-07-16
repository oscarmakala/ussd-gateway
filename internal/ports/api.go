package ports

import (
	"context"
	"ussd-gateway-go/internal/application/core/domain"
)

type APIPort interface {
	ProcessRequest(ctx context.Context, ussdRequest domain.UssdRequest) (domain.UssdResponse, error)
}
