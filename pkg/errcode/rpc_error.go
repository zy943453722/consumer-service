package errcode

import (
	pb "consumer-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//客户端使用
type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func ToRPCStatus(code int, msg string) *Status {
	s, _ := status.New(ToRPCCode(code), msg).WithDetails(&pb.Error{
		Code:    int32(code),
		Message: msg,
	})
	return &Status{s}
}

func TogRPCError(err *Error, e error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(&pb.Error{Code: int32(err.Code()), Message: e.Error()})
	return s.Err()
}

//ToRPCCode 自定义错误转RPC错误
func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	case ErrorGetConsumerListFail.Code():
		statusCode = codes.Unimplemented
	case ErrorJsonConvertFail.Code():
		statusCode = codes.Internal
	default:
		statusCode = codes.Unknown
	}

	return statusCode
}
