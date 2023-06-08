package main

import (
	"fmt"
	"phoenixbuilder/GameControl/external/encoding"
)

func main() {
	buf := encoding.Buffer{}
	buf.InitBuffer()
	buf.GetBuffer().Write([]byte{1, 0, 1})
	buf.GetBuffer().Write([]byte("ss"))
	fmt.Println(buf.GetBuffer().Bytes())
	ans, err := buf.DecodeString()
	if err != nil {
		panic(err)
	}
	fmt.Println(*ans)
}
