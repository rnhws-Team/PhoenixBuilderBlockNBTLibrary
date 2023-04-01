package main

import "fmt"

type ABC struct {
	num int32
}

func change(s ABC) {
	s.num = 524
}

func main() {
	new := ABC{}
	change(new)
	fmt.Println(new)
}
