package main

import (
	"fmt"
	"github.com/wolferhua/goomem"
)

func main() {
	goomem.StartUp()
	v, err := goomem.Get("a")
	if err != nil {
		fmt.Println(v)
		fmt.Println(err)
	}
}
