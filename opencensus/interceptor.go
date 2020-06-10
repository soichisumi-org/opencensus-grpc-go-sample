package opencensus

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/soichisumi/go-util/logger"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

var exporter *stackdriver.Exporter

func initExporter(project string) {
	if exporter != nil {
		return
	}
	_exporter, err := stackdriver.NewExporter(stackdriver.Options{ProjectID: project})
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	exporter = _exporter
	view.RegisterExporter(exporter)
	view.SetReportingPeriod(15 * time.Second) //TODO
}

func InitClientTrace(project string) {
	initExporter(project)

	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		logger.Fatal("", zap.Error(err))
	}
}

func InitServerTrace(project string) {
	initExporter(project)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		logger.Fatal("", zap.Error(err))
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func UnaryClientTraceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, span := trace.StartSpan(ctx, method)
		defer span.End()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func UnaryServerTraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
		ctx, span := trace.StartSpan(ctx, info.FullMethod)
		defer span.End()

		return handler(ctx, req)
	}
}