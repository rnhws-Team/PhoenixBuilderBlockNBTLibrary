package main

import (
	"encoding/gob"
	"fmt"
	"phoenixbuilder/GameControl/ResourcesControlCenter"
)

func main() {
	a := map[string]any{"2": map[int]any{2: 3}}
	var b map[string]any
	gob.Register(map[string]interface{}{})
	ResourcesControlCenter.DeepCopy(
		&a,
		&b,
		func() {
			gob.Register(map[string]any{})
			gob.Register(map[int]any{})
		},
	)
	fmt.Printf("%#v\n", b)
}
