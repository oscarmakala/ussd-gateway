package http

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"ussd-gateway-go/internal/ports"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *http.ServeMux
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	httpServer := http.NewServeMux()
	a.server = httpServer
	log.Printf("starting order service on port %d ...", a.port)
	if err := http.Serve(listen, a.server); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}

}
