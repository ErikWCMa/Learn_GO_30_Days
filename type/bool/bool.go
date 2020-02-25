package main

import "fmt"

func main() {
	var a bool
	a = true
	fmt.Println("a=", a)

	b := false
	fmt.Println("b=", b)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
