package encoding

import "bytes"

// 创建一个新的阅读器
func NewReader(reader *bytes.Buffer) *Reader {
	return &Reader{r: reader}
}

// 创建一个新的写入者
func NewWriter(writer *bytes.Buffer) *Writer {
	return &Writer{w: writer}
}

// 为二进制数据实现的 IO 操作流。
// 以下列出的每个函数都提供了两个实现，
// 以允许 Marshal 或 UnMarshal 二进制数据
type IO interface {
	GetBuffer() (*bytes.Buffer, bool)
	Slice(x *[]byte) error
	String(x *string) error
	Bool(x *bool) error
	Uint8(x *uint8) error
	Int8(x *int8) error
	Uint16(x *uint16) error
	Int16(x *int16) error
	Uint32(x *uint32) error
	Int32(x *int32) error
	Uint64(x *uint64) error
	Int64(x *int64) error
}
