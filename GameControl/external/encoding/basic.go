package encoding

import (
	"bytes"
	"io"
)

// 用于读取二进制切片的阅读器
type Reader struct {
	r interface {
		io.Reader
		io.ByteReader
	}
}

// 用于写入二进制切片的写入者
type Writer struct {
	w interface {
		io.Writer
		io.ByteWriter
	}
}

// 创建一个新的阅读器
func NewReader(reader *bytes.Buffer) *Reader {
	return &Reader{r: reader}
}

// 创建一个新的写入者
func NewWriter(writer *bytes.Buffer) *Writer {
	return &Writer{w: writer}
}

// 取得阅读器的底层切片
func (r *Reader) GetBuffer() (*bytes.Buffer, bool) {
	ans, err := r.r.(*bytes.Buffer)
	return ans, err
}

// 取得写入者的底层切片
func (w *Writer) GetBuffer() (*bytes.Buffer, bool) {
	ans, err := w.w.(*bytes.Buffer)
	return ans, err
}
