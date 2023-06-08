package external

import (
	"phoenixbuilder/GameControl/GlobalAPI"
)

// 指代当前引擎的数据版本。
// 此字段可能会用作后向兼容的依据
const CurrentVersion = "1.0.0"

type Value interface {
	Run(env *GlobalAPI.GlobalAPI) Return // 指代用于实际执行当次请求的函数
}               // 指代要传入到函数的参数的具体内容
type Key string // 指代要传入到函数的参数的名称

type Input map[Key]Value // 指代要向函数传入的所有参数
type Return interface{}  // 指代函数的返回值
