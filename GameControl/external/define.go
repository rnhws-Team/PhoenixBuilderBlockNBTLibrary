package external

import (
	"phoenixbuilder/GameControl/GlobalAPI"
)

// ------------------------- General -------------------------

// 指代当前引擎的数据版本。
// 此字段可能会用作后向兼容的依据
const CurrentVersion = "1.0.0"

type Value interface {
	Run(env *GlobalAPI.GlobalAPI) Return // 指代用于实际执行当次请求的函数
}               // 指代要传入到函数的参数的具体内容
type Key string // 指代要传入到函数的参数的名称

type Input map[Key]Value // 指代要向函数传入的所有参数
type Return interface{}  // 指代函数的返回值

// ------------------------- Request -------------------------

// 描述单次请求的基本信息
type RequestHeader struct {
	// 指代数据版本。此字段可能会用作后向兼容的依据
	Version string `json:"version"`
	// 指代回声，用于使用者区别每个请求的响应包。
	// 如果指定了此字段，
	// 那么响应包中也会包含一个完全相同的 request_id 字段。
	// 如果没有提供此字段，则响应包中的 request_id 将为 null
	Echo *string `json:"request_id"`
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
	Module    *string `json:"module"`         // 指代要访问的模块
	SubModule *string `json:"sub_module"`     // 指代要访问的子模块
	FuncName  *string `json:"function_name"`  // 指代要访问的函数
	FuncInput Input   `json:"function_input"` // 指代要向函数传入的参数
}

// 描述单个的请求
type Request struct {
	Header RequestHeader `json:"header"` // 描述当次请求的基本信息
	Body   RequestBody   `json:"body"`   // 指定当次请求的详细信息
}

// ------------------------- Responce -------------------------

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

// ------------------------- END -------------------------
