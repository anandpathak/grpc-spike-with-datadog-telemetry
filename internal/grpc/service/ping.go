package service

import (
	"context"
	"demogrpc/internal/grpc/domain"
)

type PingService struct{
}

func (p PingService) Ping(context context.Context,message *domain.PingMessage)  (*PingResponse, error) {
	return &PingResponse{
		Response:             &domain.PingMessage{
			Test:                 "pong",
		},
	}, nil
}
