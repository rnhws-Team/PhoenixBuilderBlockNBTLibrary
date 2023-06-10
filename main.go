package main

import (
	"bytes"
	"fmt"
	"phoenixbuilder/encoding"
)

func main() {
	writer := encoding.NewWriter(bytes.NewBuffer([]byte{}))
	var new encoding.IO
	new = writer
	num := int8(-23)
	new.Int8(&num)

	get, _ := writer.GetBuffer()
	reader := encoding.NewReader(get)
	new = reader
	newNum := int8(0)
	new.Int8(&newNum)
	fmt.Println(newNum)
}
