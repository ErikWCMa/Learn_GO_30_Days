package main

import "fmt"

func main() {
	var x [4]float64
	x[0] = 23
	x[1] = 45
	x[2] = 33
	x[3] = 21
	var total float64 = 0
	for _, value := range x {
		total = total + value
	}
	fmt.Println(total / float64(len(x)))
}
