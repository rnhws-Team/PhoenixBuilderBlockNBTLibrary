package external

// 描述单个请求所对应的响应
type Responce struct {
	// 指代此响应包所对应的请求包的回声。
	// 只有此响应包对应的请求包中提供了此字段时才会存在，
	// 否则为 null
	Echo *string `json:"request_id"`
	// 指代对应请求的完成时间
	FinishTime string `json:"finish_time"`
	// 指代请求的函数的返回值
	FuncReturn Return `json:"function_return"`
}
