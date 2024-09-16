package handler

import (
	"api-gateway/internal/service"
	"context"
)

type HandlerST struct {
	Service *service.ServiceRepositoryClient
}

func NewApiHandler(service *service.ServiceRepositoryClient) *HandlerST {
	return &HandlerST{
		Service: service,
	}
}

var ctx = context.Background()
