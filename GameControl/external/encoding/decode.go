package encoding

import (
	"encoding/binary"
	"fmt"
)

// 从阅读器阅读 length 个字节
func (r *Reader) ReadBytes(length int) ([]byte, error) {
	ans := make([]byte, length)
	_, err := r.r.Read(ans)
	if err != nil {
		return nil, fmt.Errorf("ReadBytes: %v", err)
	}
	return ans, nil
}

// 从阅读器阅读一个字符串并返回到 x 上
func (r *Reader) String(x *string) error {
	var length uint16
	err := binary.Read(r.r, binary.BigEndian, &length)
	if err != nil {
		return fmt.Errorf("(r *Reader) String: %v", err)
	}
	// get the length of the target string
	stringBytes, err := r.ReadBytes(int(length))
	if err != nil {
		return fmt.Errorf("(r *Reader) String: %v", err)
	}
	*x = string(stringBytes)
	// get the target string
	return nil
	// return
}
