package main

import "fmt"

var z string = "Suh"

type hotdog string

func main() {
	fmt.Println(foo())
	x := 42
	fmt.Println(x)
	var y = 47
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	var b hotdog = "Suh"
	fmt.Printf("%v, %T", b, b)
}

func foo() string {
	return `"Shiiiii"` + " Kidney"
}
