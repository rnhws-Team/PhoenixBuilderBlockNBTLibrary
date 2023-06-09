package main

import (
	"bytes"
	"fmt"
	"phoenixbuilder/GameControl/external/encoding"
)

func main() {
	reader := encoding.NewReader(bytes.NewBuffer([]byte{0, 1, 98}))
	m := ""
	encoding.IO.String(reader, &m)
	fmt.Println(m)

	writer := encoding.NewWriter(bytes.NewBuffer([]byte{}))
	m = "b"
	encoding.IO.String(writer, &m)
	got, _ := writer.GetBuffer()
	fmt.Println(got.Bytes())
}
