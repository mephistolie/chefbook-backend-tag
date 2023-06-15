package grpc

import (
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-tag/internal/transport/dependencies/service"
)

type TagServer struct {
	api.UnsafeTagServiceServer
	service service.Service
}

func NewServer(service service.Service) *TagServer {
	return &TagServer{service: service}
}
