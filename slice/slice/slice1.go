package main

import "fmt"

func main() {
	array1 := make([]float64, 5)
	fmt.Println(len(array1), cap(array1))

	array2 := make([]float64, 5, 10)
	fmt.Println(len(array2), cap(array2))

	array3 := [5]float64{1, 2, 3, 4, 5}
	array4 := array3[:]
	array5 := array3[0:1]
	array6 := array3[1:4]
	array7 := array3[4:5]
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(array6)
	fmt.Println(array7)

}
