package opencensus

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/soichisumi/go-util/logger"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func SetupExporter(project string) *stackdriver.Exporter {
	exporter, err := stackdriver.NewExporter(
		stackdriver.Options{
			ProjectID: project,
			//MonitoredResource: monitoredresource.Autodetect(),
			ReportingInterval: 1 * time.Minute,
			//DefaultTraceAttributes: make(map[string]interface{}),
		},
	)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	view.RegisterExporter(exporter)
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	return exporter
}

func InitTrace() {
	//if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
	//	logger.Fatal("", zap.Error(err))
	//}
	//if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
	//	logger.Fatal("", zap.Error(err))
	//}
}

func UnaryClientTraceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, span := trace.StartSpan(ctx, method)
		defer span.End()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func UnaryServerTraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, span := trace.StartSpan(ctx, info.FullMethod)
		defer span.End()

		return handler(ctx, req)
	}
}
