package errcode

import (
	pb "consumer-service/proto"
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type httpError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func GrpcGatewayError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpErr := httpError{
		Code:    int32(s.Code()),
		Message: s.Message(),
	}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpErr.Code = v.Code
			httpErr.Message = s.Message() + " caused by " + v.Message
		}
	}

	resp, _ := json.Marshal(httpErr)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	w.Write(resp)
}
