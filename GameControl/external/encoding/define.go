package encoding

import "bytes"

const StringLengthMaxLimited = 65535 // 单个字符串的最大长度上限

// 用于读写二进制数据包的结构体
type Buffer struct {
	buffer *bytes.Buffer
}

// 初始化 Buffer 为可读写状态。
// 每次调用此函数后，底层切片将会被重置为空
func (b *Buffer) InitBuffer() {
	b.buffer = bytes.NewBuffer([]byte{})
}

// 将底层切片替换为 buffer
func (b *Buffer) ReplaceBuffer(buffer *bytes.Buffer) {
	b.buffer = buffer
}

// 取得 Buffer 的底层切片，是一个指针
func (b *Buffer) GetBuffer() *bytes.Buffer {
	return b.buffer
}
