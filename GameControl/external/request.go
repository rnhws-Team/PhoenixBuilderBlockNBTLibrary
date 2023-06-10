package external

import (
	"phoenixbuilder/encoding"
)

// 描述单次请求的基本信息
type RequestHeader struct {
	// 指代数据版本。此字段可能会用作后向兼容的依据
	Version string `json:"version"`
	// 指代回声，用于使用者区别每个请求的响应包。
	// 如果指定了此字段，
	// 那么响应包中也会包含一个完全相同的 request_id 字段。
	Echo string `json:"request_id"`
	// 指代是否返回当次请求对应的响应包。
	// 当为假时，请求结果将不会送回，
	// 也不会存在回声
	GetResponce bool `json:"get_responce"`
	// 指代请求者，这并不是必须的
	Requester string `json:"requester"`
	// 是否在控制台打印函数的执行状况
	PrintRunningSituation bool `json:"print_running_situation"`
	// 是否在抛出惊慌时抑制错误
	SuppressError bool `json:"suppress_error"`
}

// 描述单次请求的详细信息
type RequestBody struct {
	Module    string `json:"module"`         // 指代要访问的模块
	SubModule string `json:"sub_module"`     // 指代要访问的子模块
	FuncName  string `json:"function_name"`  // 指代要访问的函数
	FuncInput Input  `json:"function_input"` // 指代要向函数传入的参数
}

// 描述单个的请求
type Request struct {
	Header RequestHeader `json:"header"` // 描述当次请求的基本信息
	Body   RequestBody   `json:"body"`   // 指定当次请求的详细信息
}

// AutoMarshal...
func (r *Request) Marshal(io encoding.IO) {
	TestError(io.String(&r.Header.Version))
	TestError(io.String(&r.Header.Echo))
	TestError(io.Bool(&r.Header.GetResponce))
	TestError(io.String(&r.Header.Requester))
	TestError(io.Bool(&r.Header.PrintRunningSituation))
	TestError(io.Bool(&r.Header.SuppressError))
	// header
	TestError(io.String(&r.Body.Module))
	TestError(io.String(&r.Body.SubModule))
	TestError(io.String(&r.Body.FuncName))
	r.Body.FuncInput.Marshal(io)
	// body
}
