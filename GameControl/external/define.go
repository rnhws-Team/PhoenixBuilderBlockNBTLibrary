package external

import (
	"phoenixbuilder/GameControl/GlobalAPI"
	"phoenixbuilder/encoding"
)

// 指代当前引擎的数据版本。
// 此字段可能会用作后向兼容的依据
const CurrentVersion = "DEV"

// 指代要向函数传入的所有参数
type Input interface {
	// 用于执行当次请求的实现
	Run(env *GlobalAPI.GlobalAPI) Return
	// 当传入 encoding.Reader 时，
	// 数据将从 encoding.Reader 解码至 Input ；
	// 当传入 encoding.Writer 时，
	// Input 将被编码至 encoding.Writer
	Marshal(io encoding.IO)
}

// 指代函数的返回值
type Return interface {
	// 当传入 encoding.Reader 时，
	// 数据将从 encoding.Reader 解码至 Return ；
	// 当传入 encoding.Writer 时，
	// Return 将被编码至 encoding.Writer
	Marshal(io encoding.IO)
}
