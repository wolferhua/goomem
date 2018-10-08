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
	goomem.Set("a", User{Age: 100}, 100)

	v, err := goomem.Get("a")
	i := v.(User)

	//var i User1
	//goomem.GetForType("a", &i)
	fmt.Println(i.Age)
	fmt.Println(i.Age)
	fmt.Println(i.Age)

	i.Age = 1000
	//var b User1
	//goomem.GetForType("a", &b)
	//fmt.Println(b.Age)
	//fmt.Println(b.Age)
	//fmt.Println(b.Age)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	fmt.Println(i)
	//fmt.Println(b)
}
