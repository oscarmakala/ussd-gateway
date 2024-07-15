package ports

import "ussd-gateway-go/internal/application/core/domain"

type StoragePort interface {
	LoadNode(moduleName string, applicationId string) (domain.Node, error)
	LoadProjectOptions(applicationId string) (domain.ProjectIndex, error)
}
