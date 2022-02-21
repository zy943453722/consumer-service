package middleware

import (
	"consumer-service/global"
	"consumer-service/pkg/metatext"
	"context"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/opentracing/opentracing-go/ext"

	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
)

func defaultContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc

	if _, ok := ctx.Deadline(); !ok {
		defaultTime := 10 * time.Second
		ctx, cancel = context.WithTimeout(ctx, defaultTime)
	}

	return ctx, cancel
}

func ContextTimeout(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, cancel := defaultContextTimeout(ctx)
	if cancel != nil {
		defer cancel()
	}

	return invoker(ctx, method, req, resp, cc, opts...)
}

func ClientTracing(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var (
		parentCtx opentracing.SpanContext
		spanOpts  []opentracing.StartSpanOption
	)

	//从结构activeSpanKey中获取span的信息作为父级的span
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		//从父级span中分析出父级的context
		parentCtx = parentSpan.Context()
		spanOpts = append(spanOpts, opentracing.ChildOf(parentCtx))
	}
	spanOpts = append(spanOpts, []opentracing.StartSpanOption{
		opentracing.Tag{
			Key:   string(ext.Component),
			Value: "gRPC",
		},
		ext.SpanKindRPCClient,
	}...)

	span := global.Tracer.StartSpan(method, spanOpts...)
	defer span.Finish()

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	global.Tracer.Inject(span.Context(), opentracing.TextMap, metatext.MetadataTextMap{MD: md})
	newCtx := opentracing.ContextWithSpan(metadata.NewOutgoingContext(ctx, md), span)
	return invoker(newCtx, method, req, reply, cc, opts...)
}
