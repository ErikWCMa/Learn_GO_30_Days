package main

import "fmt"

func main() {
	slices1 := []int{1, 2, 3, 4, 5}
	slices2 := make([]int, 2, 3)
	copy(slices2, slices1)
	fmt.Println(slices1, slices2)
	fmt.Println(len(slices2), cap(slices2))
	slices2 = append(slices2, 5)
	//slices2[2] = 6 error
	fmt.Println(slices2[2], len(slices2), cap(slices2))
	slices2[2] = 6
	fmt.Println(slices2[2])
}
