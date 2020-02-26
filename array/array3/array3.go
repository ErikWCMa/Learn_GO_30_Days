package main

import "fmt"

func main() {
	var x [4]float64
	x[0] = 23
	x[1] = 45
	x[2] = 33
	x[3] = 21
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total = total + x[i]
	}
	fmt.Println(total / float64(len(x)))
}
