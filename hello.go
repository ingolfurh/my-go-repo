package main

import "fmt"

type pulsa int

var x pulsa
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 4
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
