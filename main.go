package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(add(3, 5))
}

func add(a, b int) (sum int) {
	return a + b
}
