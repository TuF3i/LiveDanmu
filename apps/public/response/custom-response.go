package response

var (
	OperationSuccess        = Response{Status: 20000, Info: "Operation Success"}           // OperationSuccess 执行成功
	EmptyJWTString          = Response{Status: 40001, Info: "Empty JWT String"}            // EmptyJWTString 空JWT字符串
	JWTNotRegisteredInRedis = Response{Status: 40002, Info: "JWT Not Registered In Redis"} // JWTNotRegisteredInRedis JWT未注册或已过期
	ValidateRequestFail     = Response{Status: 40003, Info: "Validate Request Fail"}       // ValidateRequestFail 校验请求失败
	EmptyRVID               = Response{Status: 40004, Info: "Empty RVID"}                  // EmptyRVID 空RVID
)
