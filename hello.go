package main

import "fmt"

func main() {
	fmt.Println("Suh dudes")
	for i := 0; i < 10; i++ {
		fmt.Println(foo())
	}
}

func foo() string {
	return "Shiiiii"
}
