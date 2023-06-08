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
