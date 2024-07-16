package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
	"ussd-gateway-go/internal/ports"
)

type Adapter struct {
	api  ports.APIPort
	port int
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

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	log.Printf("starting ussd gateway service on port %d ...", a.port)
	if err := http.Serve(listen, r); err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
}
