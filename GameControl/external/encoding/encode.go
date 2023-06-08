package encoding

import (
	"encoding/binary"
	"fmt"
)

// 向 Buffer 写入字节切片 bytes
func (b *Buffer) WriteBytes(bytes []byte) error {
	_, err := b.buffer.Write(bytes)
	if err != nil {
		return fmt.Errorf("WriteBytes: %v", err)
	}
	return nil
}

// 向 Buffer 写入一个数据头。
// 如果当前数据不存在，也就是 value 为 nil ，
// 则写入 0 ，否则写入 1 。
// 如果写入的数据头为 0 ，则返回假，否则返回真
func (b *Buffer) EncodeHeader(value interface{}) bool {
	if value == nil {
		b.buffer.Write([]byte{0})
		return false
	} else {
		b.buffer.Write([]byte{1})
		return true
	}
}

// 向 Buffer 写入字符串 str 。
// 如果 str 为 nil ，则仅会写入 0 ，
// 否则写入 1 并写入字符串数据
func (b *Buffer) EncodeString(str *string) error {
	if len(*str) > 65535 {
		return fmt.Errorf("EncodeString: The length of the target string is out of the max limited %v", StringLengthMaxLimited)
	}
	// check length
	if !b.EncodeHeader(str) {
		return nil
	}
	// write header
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, uint16(len(*str)))
	b.buffer.Write(lengthBytes)
	// write the length of the target string
	err := b.WriteBytes([]byte(*str))
	if err != nil {
		return fmt.Errorf("EncodeString: %v", err)
	}
	// write string
	return nil
}
