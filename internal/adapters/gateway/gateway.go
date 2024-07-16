package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"ussd-gateway-go/gen/go/proto/ussd/v1"
	"ussd-gateway-go/internal/ports"
)

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Network, Addr string
}

type Adapter struct {
	api  ports.APIPort
	port int
	ussd.UnimplementedUssdServiceServer
	server *grpc.Server
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	a.server = grpcServer

	ussd.RegisterUssdServiceServer(grpcServer, a)
	ctx := context.Background()

	mux := http.NewServeMux()

	gateway := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := ussd.RegisterUssdServiceHandlerFromEndpoint(ctx, gateway, fmt.Sprintf("0.0.0.0:%d", a.port), opts)
	if err != nil {
		log.Fatal("Failed to register gateway handler: ", err)
	}

	mux.Handle("/", gateway)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	server := http.Server{
		Addr: fmt.Sprintf(":%d", a.port),
	}

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}
