package main

import (
	"bytes"
	"fmt"
	"phoenixbuilder/lib/encoding"

	"github.com/google/uuid"
)

func main() {
	writer := encoding.NewWriter(bytes.NewBuffer([]byte{}))
	var new encoding.IO
	new = writer
	uniqueID, _ := uuid.NewUUID()
	new.UUID(&uniqueID)

	get, _ := writer.GetBuffer()
	reader := encoding.NewReader(get)
	new = reader
	newUniqueID := uuid.UUID{}
	new.UUID(&newUniqueID)
	fmt.Println(newUniqueID.String())
}
