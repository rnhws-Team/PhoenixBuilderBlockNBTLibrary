package encoding

// 为 external 二进制数据包实现的 IO 操作流。
// 以下列出的每个函数都提供了两个实现，
// 以允许 Marshal 或 UnMarshal 二进制数据
type IO interface {
	String(x *string) error
}
