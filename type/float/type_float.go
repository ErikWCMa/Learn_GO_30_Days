package main

import "fmt"

func main() {
	var floatValue float64
	floatValue = 7.0
	var floatValue2 = 3.0
	fmt.Println("7.0/3.0 =", floatValue/floatValue2)

	var test float64
	var test2 float32
	test = 1.1
	test2 = 2.2
	fmt.Println("test+test2 =", float32(test)+test2)
}
