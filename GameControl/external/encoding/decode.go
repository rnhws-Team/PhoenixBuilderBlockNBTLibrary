package encoding

import (
	"encoding/binary"
	"fmt"
)

// 从 Buffer 阅读 length 个字节。
// 如果阅读过程中发生了错误，
// 亦或阅读所得的字节数少于 length ，
// 那么此函数将返回错误
func (b *Buffer) ReadBytes(length int) ([]byte, error) {
	ans := make([]byte, length)
	_, err := b.buffer.Read(ans)
	if err != nil {
		return nil, fmt.Errorf("ReadBytes: %v", err)
	}
	return ans, nil
}

// 从 Buffer 阅读一个数据头。
// 如果当前数据不存在，也就是其应当为 nil ，
// 则返回值第一项为假，否则为真
func (b *Buffer) DecodeHeader() (bool, error) {
	headerBytes, err := b.ReadBytes(1)
	if err != nil {
		return false, fmt.Errorf("DecodeReader: %v", err)
	}
	// get header
	switch header := headerBytes[0]; header {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("DecodeHeader: Unexpected header %#v was find", header)
	}
	// check header and return
}

// 从 buffer 阅读一个字符串
func (b *Buffer) DecodeString() (*string, error) {
	exist, err := b.DecodeHeader()
	if err != nil {
		return nil, fmt.Errorf("DecodeString: %v", err)
	}
	if !exist {
		return nil, nil
	}
	// check if this is exist
	var length uint16
	err = binary.Read(b.buffer, binary.BigEndian, &length)
	if err != nil {
		return nil, fmt.Errorf("DecodeString: %v", err)
	}
	// get the length of the target string
	stringBytes, err := b.ReadBytes(int(length))
	if err != nil {
		return nil, fmt.Errorf("DecodeString: %v", err)
	}
	str := string(stringBytes)
	// get the target string
	return &str, nil
	// return
}
