package errcode

var (
	ErrorGetConsumerListFail = NewError(30001, "获取客户列表失败")
	ErrorJsonConvertFail     = NewError(30002, "json转化失败")
)
