package grpc

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ussd-gateway-go/gen/go/proto/ussd/v1"
	"ussd-gateway-go/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *ussd.UssdRequest) (*ussd.UssdResponse, error) {
	log.WithContext(ctx).Info("Creating ussd request...")
	newUssdRequest := domain.NewUssdRequest(
		request.UssdPayload,
		request.Language,
		request.SessionId,
		request.Msisdn,
	)
	result, err := a.api.ProcessRequest(ctx, newUssdRequest)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &ussd.UssdResponse{
		UssdPayload:     result.Payload,
		UssdMessageType: "",
		ServiceCode:     "",
		SessionId:       "",
		Msisdn:          "",
	}, nil
}
