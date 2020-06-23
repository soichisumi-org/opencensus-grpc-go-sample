package main

import (
	"context"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"math/rand"
)

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

type EchoServer struct{}

func (EchoServer) Echo(ctx context.Context, req *grpctesting.EchoRequest) (*grpctesting.EchoResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	res := ""
	for i:=0; i<rand.Intn(10); i++{
		_ctx, span := trace.StartSpan(ctx, "childserver.Echo.loop")
		res = res + " " +req.Message
		span.End()
		ctx = _ctx
	}
	logger.Info("", zap.Any("req", req), zap.Any("md", md))
	return &grpctesting.EchoResponse{
		Message:              res,
	}, nil
}

func (EchoServer) Empty(ctx context.Context, req *grpctesting.EmptyRequest) (*grpctesting.EmptyResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	ctx, span := trace.StartSpan(ctx, "childserver.Empty.loop")
	logger.Info("", zap.Any("req", req), zap.Any("md", md))
	span.End()
	return &grpctesting.EmptyResponse{}, nil
}
