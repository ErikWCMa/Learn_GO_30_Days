package main

import "fmt"

func main() {
	slices1 := []int{1, 2, 3}
	slices2 := append(slices1, 4, 5)
	fmt.Println(slices1, slices2)
}
