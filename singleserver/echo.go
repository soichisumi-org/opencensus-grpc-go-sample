package main

import (
	"context"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

type EchoServer struct {}

func (e *EchoServer) Echo(ctx context.Context, req *grpctesting.EchoRequest) (*grpctesting.EchoResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	ctx, span := trace.StartSpan(ctx, "parentserver.Echo.log")
	logger.Info("", zap.Any("req", req), zap.Any("md", md))
	span.End()
	return &grpctesting.EchoResponse{Message: req.Message}, nil
}

func (e *EchoServer) Empty(ctx context.Context, req *grpctesting.EmptyRequest) (*grpctesting.EmptyResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	ctx, span := trace.StartSpan(ctx, "parentserver.Empty.log")
	logger.Info("", zap.Any("req", req), zap.Any("md", md))
	span.End()
	return &grpctesting.EmptyResponse{}, nil
}
