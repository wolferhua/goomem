package main

import (
	"fmt"
	"github.com/wolferhua/goomem/goomem"
)


func main() {
	fmt.Println(goomem.Get("a"))
	fmt.Println(goomem.Set("a", "a", 0))
	fmt.Println(goomem.Get("a"))
	fmt.Println(goomem.Set("a", 1231, 0))
	fmt.Println(goomem.Get("a"))
	fmt.Println(goomem.Set("a", 111, 2))
	fmt.Println(goomem.Get("a"))
	fmt.Println(goomem.Set("a", 123123, 3))
	fmt.Println(goomem.Get("a"))
	select {

	}
}
