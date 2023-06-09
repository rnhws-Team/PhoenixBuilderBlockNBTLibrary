package encoding

import (
	"encoding/binary"
	"fmt"
)

// 向写入者写入字节切片 p
func (w *Writer) WriteBytes(p []byte) error {
	_, err := w.w.Write(p)
	if err != nil {
		return fmt.Errorf("WriteBytes: %v", err)
	}
	return nil
}

// 向写入者写入字符串 x
func (w *Writer) String(x *string) error {
	if len(*x) > StringLengthMaxLimited {
		return fmt.Errorf("(w *Writer) String: The length of the target string is out of the max limited %v", StringLengthMaxLimited)
	}
	// check length
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, uint16(len(*x)))
	err := w.WriteBytes(lengthBytes)
	if err != nil {
		return fmt.Errorf("(w *Writer) String: %v", err)
	}
	// write the length of the target string
	err = w.WriteBytes([]byte(*x))
	if err != nil {
		return fmt.Errorf("(w *Writer) String: %v", err)
	}
	// write string
	return nil
}
