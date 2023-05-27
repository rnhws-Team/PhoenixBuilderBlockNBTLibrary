package ResourcesControlCenter

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 将 source 深拷贝到 destination
func DeepCopy(source interface{}, destination interface{}) error {
	gob.Register(map[string]interface{}{})
	var buffer bytes.Buffer
	// init values
	err := gob.NewEncoder(&buffer).Encode(source)
	if err != nil {
		return fmt.Errorf("DeepCopy: %v", err)
	}
	// encode
	err = gob.NewDecoder(&buffer).Decode(destination)
	if err != nil {
		return fmt.Errorf("DeepCopy: %v", err)
	}
	// decode
	return nil
	// return
}
