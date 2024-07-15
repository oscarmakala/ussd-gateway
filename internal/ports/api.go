package ports

import (
	"context"
	"ussd-gateway-go/internal/application/core/domain"
)

type APIPort interface {
	HandleUssdRequest(ctx context.Context, ussdRequest domain.UssdRequest)
}
