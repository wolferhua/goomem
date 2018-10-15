package main

import (
	"fmt"
	"reflect"
)

type intt int
type inttt intt

func main() {
	fmt.Println("fmt.Println(reflect.TypeOf(1))")
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1).Kind())

	fmt.Println("fmt.Println(reflect.TypeOf(intt(1)))")
	fmt.Println(reflect.TypeOf(intt(1)))
	fmt.Println(reflect.TypeOf(intt(1)).Kind())
	fmt.Println(reflect.TypeOf(intt(1)).IsVariadic())

	fmt.Println("fmt.Println(reflect.TypeOf(inttt(1)))")
	fmt.Println(reflect.TypeOf(inttt(1)))
	fmt.Println(reflect.TypeOf(inttt(1)).Kind())
	fmt.Println(reflect.TypeOf(inttt(1)).IsVariadic())

	fmt.Println("fmt.Println(reflect.TypeOf(1.1))")
	fmt.Println(reflect.TypeOf(1.1))
	fmt.Println(reflect.TypeOf(1.1).Kind())
	fmt.Println(reflect.TypeOf(1.1).IsVariadic())

	fmt.Println("fmt.Println(reflect.TypeOf(\"\"))")

	fmt.Println(reflect.TypeOf(string("")))
	fmt.Println(reflect.TypeOf(string("")).Kind())
	fmt.Println(reflect.TypeOf(string("")).IsVariadic())

}
