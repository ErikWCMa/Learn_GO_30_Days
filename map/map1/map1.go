package main

import "fmt"

func main() {
	x := make(map[int]int)
	x[1] = 10
	x[2] = 20
	fmt.Println(x)
	delete(x, 1)
	fmt.Println(x)

	y := make(map[string]int)
	y["key"] = 10
	fmt.Println(y)
	delete(y, "key")
	fmt.Println(y)
}
