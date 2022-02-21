package main

import (
	"consumer-service/global"
	"consumer-service/internal/auth"
	"consumer-service/internal/middleware"
	"consumer-service/pkg/errcode"
	"consumer-service/pkg/tracer"
	pb "consumer-service/proto"
	"context"
	"fmt"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func init() {
	if err := setupTracer(); err != nil {
		panic("启动tracer失败")
	}
}

func main() {
	ctx := context.Background()

	//设置grpc连接配置，包含中间件等
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient( //链式中间件
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(2), //重试次数最大2
				grpc_retry.WithCodes( //针对特定错误码进行重试
					codes.Unknown,
					codes.Internal,
					codes.DeadlineExceeded,
					codes.Unimplemented,
				),
			),
			middleware.ContextTimeout, //超时客户端主动断开连接
			middleware.ClientTracing,  //客户端的链路追踪
		),
	),
		grpc.WithPerRPCCredentials(&auth.Auth{
			AppKey:    "hello",
			AppSecret: "world",
		}), //grpc默认的token认证校验，用metadata的形式传递给服务端
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	//metadata处理,数据包传递存放在http的header，存储在context中传递
	md := metadata.New(map[string]string{"go": "programming", "tour": "book"})
	newCtx := metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.DialContext(newCtx, "localhost:8081", opts...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	c := pb.NewConsumerServiceClient(conn)
	resp, err := c.GetConsumerList(ctx, &pb.GetConsumerListReq{
		Source: "cps",
	})
	if err != nil {
		sts := errcode.FromError(err)
		details := sts.Details()
		fmt.Printf("detail is %v", details)
	}
	fmt.Printf("resp:%v", resp)
}

func setupTracer() error {
	var err error
	jagerTracer, _, err := tracer.NewJaegerTracer("article-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jagerTracer
	return nil
}
