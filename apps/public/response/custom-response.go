package response

var (
	OperationSuccess = Response{Status: 20000, Info: "Operation Success"} // 执行成功

	EmptyJWTString = Response{Status: 40001, Info: ""}
)
