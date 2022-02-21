package server

import (
	"consumer-service/pkg/errcode"
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type Auth struct{}

func (a *Auth) GetAppKey() string {
	return "hello"
}

func (a *Auth) GetAppSecret() string {
	return "world"
}

func (a *Auth) Check(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	var appKey, appSecret string
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return errcode.TogRPCError(errcode.Unauthorized, fmt.Errorf("%s", "校验不通过"))
	}

	return nil
}
