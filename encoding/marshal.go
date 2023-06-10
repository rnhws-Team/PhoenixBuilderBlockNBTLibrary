package encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 取得写入者的底层切片
func (w *Writer) GetBuffer() (*bytes.Buffer, bool) {
	ans, err := w.w.(*bytes.Buffer)
	return ans, err
}

// 向写入者写入字节切片 p
func (w *Writer) WriteBytes(p []byte) error {
	_, err := w.w.Write(p)
	if err != nil {
		return fmt.Errorf("WriteBytes: %v", err)
	}
	return nil
}

// 向写入者写入二进制切片 x
func (w *Writer) Slice(x *[]byte) error {
	if len(*x) > SliceLengthMaxLimited {
		return fmt.Errorf("(w *Writer) Slice: The length of the target slice is out of the max limited %v", SliceLengthMaxLimited)
	}
	// check length
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(len(*x)))
	err := w.WriteBytes(lengthBytes)
	if err != nil {
		return fmt.Errorf("(w *Writer) Slice: %v", err)
	}
	// write the length of the target slice
	err = w.WriteBytes(*x)
	if err != nil {
		return fmt.Errorf("(w *Writer) Slice: %v", err)
	}
	// write slice
	return nil
	// return
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
	// return
}

// 向写入者写入布尔值 x
func (w *Writer) Bool(x *bool) error {
	if *x {
		err := w.WriteBytes([]byte{1})
		if err != nil {
			return fmt.Errorf("(w *Writer) Bool: %v", err)
		}
	} else {
		err := w.WriteBytes([]byte{0})
		if err != nil {
			return fmt.Errorf("(w *Writer) Bool: %v", err)
		}
	}
	return nil
}

// 向写入者写入 x(uint8)
func (w *Writer) Uint8(x *uint8) error {
	err := w.WriteBytes([]byte{*x})
	if err != nil {
		return fmt.Errorf("(w *Writer) Uint8: %v", err)
	}
	return nil
}

// 向写入者写入 x(int8)
func (w *Writer) Int8(x *int8) error {
	err := w.WriteBytes([]byte{byte(*x)})
	if err != nil {
		return fmt.Errorf("(w *Writer) Int8: %v", err)
	}
	return nil
}

// 向写入者写入 x(uint16)
func (w *Writer) Uint16(x *uint16) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, *x)
	err := w.WriteBytes(tmp)
	if err != nil {
		return fmt.Errorf("(w *Writer) Uint16: %v", err)
	}
	return nil
}

// 向写入者写入 x(int16)
func (w *Writer) Int16(x *int16) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, uint16(*x))
	err := w.WriteBytes(tmp)
	if err != nil {
		return fmt.Errorf("(w *Writer) Int16: %v", err)
	}
	return nil
}

// 向写入者写入 x(uint32)
func (w *Writer) Uint32(x *uint32) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint32(tmp, *x)
	err := w.WriteBytes(tmp)
	if err != nil {
		return fmt.Errorf("(w *Writer) Uint32: %v", err)
	}
	return nil
}

// 向写入者写入 x(int32)
func (w *Writer) Int32(x *int32) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint32(tmp, uint32(*x))
	err := w.WriteBytes(tmp)
	if err != nil {
		return fmt.Errorf("(w *Writer) Int32: %v", err)
	}
	return nil
}

// 向写入者写入 x(uint64)
func (w *Writer) Uint64(x *uint64) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint64(tmp, *x)
	err := w.WriteBytes(tmp)
	if err != nil {
		return fmt.Errorf("(w *Writer) Uint64: %v", err)
	}
	return nil
}

// 向写入者写入 x(int64)
func (w *Writer) Int64(x *int64) error {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint64(tmp, uint64(*x))
	err := w.WriteBytes(tmp)
	if err != nil {
		return fmt.Errorf("(w *Writer) Int64: %v", err)
	}
	return nil
}
