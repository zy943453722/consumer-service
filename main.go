package main

import (
	"consumer-service/global"
	"consumer-service/internal/middleware"
	"consumer-service/pkg/errcode"
	"consumer-service/pkg/tracer"
	pb "consumer-service/proto"
	"consumer-service/server"
	"context"
	"log"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func init() {
	if err := setupTracer(); err != nil {
		panic("启动tracer失败")
	}
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("consumer-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func main() {
	httpServer := runHttpServer()
	grpcServer := runGrpcServer()
	gwServer := runGrpcGatewayServer()

	//所有http请求均转发到网关
	httpServer.Handle("/", gwServer)
	/*listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("listen err:%v", err)
	}

	if err = s.Serve(listen); err != nil {
		log.Fatalf("start server err:%v", err)
	}*/
	if err := http.ListenAndServe("localhost:8081", grpcHandlerFunc(grpcServer, httpServer)); err != nil {
		log.Fatalf("start server err:%v", err)
	}
}

func runHttpServer() *http.ServeMux {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return serverMux
}

func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
			middleware.ServerTracing,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterConsumerServiceServer(s, server.NewConsumerServer())
	//使用grpcurl调试时，需要注册反射服务
	reflection.Register(s)

	return s
}

//runGrpcGatewayServer grpc网关服务，可对外提供http请求，网关作为client端，rpc请求server端
func runGrpcGatewayServer() *runtime.ServeMux {
	//endpoint指grpc server集群地址
	endpoint := "localhost:8081"
	runtime.HTTPError = errcode.GrpcGatewayError
	gwmux := runtime.NewServeMux()
	//注册http
	pb.RegisterConsumerServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})

	return gwmux
}

//grpcHandlerFunc grpc服务和http服务使用同一端口，需针对grpc和http分流
//分流规则：
//1、对 ProtoMajor 进行判断，该字段代表客户端请求的版本号，客户端始终使用 HTTP/1.1 或 HTTP/2
//2、对http Header头判断,Content-Type对应grpc是"application/grpc"
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	//h2c是明文运行HTTP/2协议,此handler拦截所有h2c流量，并将不同流量重定向到对应handler处理
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
