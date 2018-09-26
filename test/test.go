package main

import (
	"fmt"
	"github.com/wolferhua/goomem"
)

type User struct {
	Age int
}
type User1 struct {
	Age int
}

func main() {
	goomem.StartUp()
	goomem.Set("a", new(User), 100)

	v, err := goomem.Get("a")
	var i User1
	goomem.GetForType("a", &i)
	fmt.Println(i.Age)
	fmt.Println(i.Age)
	fmt.Println(i.Age)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
