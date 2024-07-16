module ussd-gateway-go

go 1.21

toolchain go1.23rc1

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/sirupsen/logrus v1.9.3
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.28.0
	go.opentelemetry.io/otel/trace v1.28.0
	google.golang.org/grpc v1.65.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.4.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
