package main

import "fmt"

func main() {
	arr := []float64{1, 2, 3, 4, 5}
	arr2 := arr[2:4]

	fmt.Println("arr", arr)
	fmt.Println("arr2", arr2)
	arr = append(arr, 6)
	fmt.Println("\nChange Origin Data\n")
	fmt.Println("arr", arr)
	fmt.Println("arr2", arr2)

	fmt.Println("\nAnother Case")

	arr3 := []float64{1, 2, 3, 4, 5}
	arr4 := append(arr3, 6)
	fmt.Println("\nChange Origin Data\n")
	arr3 = append(arr3, 7)

	fmt.Println("arr3", arr3)
	fmt.Println("arr4", arr4)
}