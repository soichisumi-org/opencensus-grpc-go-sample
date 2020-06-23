package opencensus

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"encoding/json"
	"github.com/soichisumi/go-util/logger"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math"
	"strconv"
	"strings"
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
	//view.RegisterExporter(exporter)
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

func toAttributes(req interface{}) []trace.Attribute {
	var m map[string]interface{}
	i, err := json.Marshal(req)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil
	}
	err = json.Unmarshal(i, &m)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil
	}

	var res = make([]trace.Attribute, 0, len(m))
	for k, v := range m {
		switch _v := v.(type) {
		case bool:
			res = append(res, trace.BoolAttribute(k, _v))
		case int32:
			res = append(res, trace.Int64Attribute(k, int64(_v)))
		case int64:
			res = append(res, trace.Int64Attribute(k, _v))
		case uint32:
			res = append(res, trace.Int64Attribute(k, int64(_v)))
		case uint64:
			if math.MaxInt64 < _v { // todo: use string attributes
				_v = math.MaxInt64
			}
			res = append(res, trace.Int64Attribute(k, int64(_v)))
		case float32:
			res = append(res, trace.StringAttribute(k, strconv.FormatFloat(float64(_v), 'f', -1, 64)))
		case float64:
			res = append(res, trace.StringAttribute(k, strconv.FormatFloat(_v, 'f', -1, 64)))
		case string:
			res = append(res, trace.StringAttribute(k, _v))
		case []byte:
			res = append(res, trace.StringAttribute(k, string(_v))) // may should be converted to hex
		default:
			//todo
		}
	}
	logger.Info("", zap.Any("attr", res))
	return res
}

func formatHeader(arr []string) string {
	switch len(arr) {
	case 0:
		return ""
	case 1:
		return arr[0]
	default:
		return "[" + strings.Join(arr, ",") + "]"
	}
}

func attributesFromContext(ctx context.Context) []trace.Attribute {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	var res []trace.Attribute
	for k, v := range md {
		res = append(res, trace.StringAttribute(k, formatHeader(v))) // todo: mask
	}
	return res
}

func UnaryClientTraceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, span := trace.StartSpan(ctx, method)
		defer span.End()
		//logger.Info("", zap.Any("req", req), zap.Any("map", req.(map[string]interface{})))
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func UnaryServerTraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, span := trace.StartSpan(ctx, info.FullMethod)
		defer span.End()
		span.AddAttributes(toAttributes(req)...)
		logger.Info("", zap.Any("req", req))
		return handler(ctx, req)
	}
}
