package main

import "fmt"

func main() {
	x := [4]float64{23, 45, 33, 21}
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total = total + x[i]
	}
	fmt.Println(total / float64(len(x)))
}
