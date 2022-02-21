package errcode

var (
	Success          = NewError(10000, "成功")
	Fail             = NewError(20001, "内部错误")
	InvalidParams    = NewError(20002, "无效参数")
	Unauthorized     = NewError(20003, "认证错误")
	NotFound         = NewError(20004, "没有找到")
	Unknown          = NewError(20005, "未知")
	DeadlineExceeded = NewError(20006, "超出最后截止期限")
	AccessDenied     = NewError(20007, "访问被拒绝")
	LimitExceed      = NewError(20008, "访问限制")
	MethodNotAllowed = NewError(20009, "不支持该方法")
)
