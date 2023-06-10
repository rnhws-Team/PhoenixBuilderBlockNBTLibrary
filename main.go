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
	MAP := map[string][]byte{"你好": {2, 0, 1, 8}, "hi": {8, 1, 0, 2}}
	new.Map(&MAP)

	get, _ := writer.GetBuffer()
	reader := encoding.NewReader(get)
	new = reader
	newMAP := map[string][]byte{}
	new.Map(&newMAP)
	fmt.Println(newMAP)
}
