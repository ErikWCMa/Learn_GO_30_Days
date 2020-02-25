package main

import "fmt"

func main() {
	var complexValue complex64
	complexValue = 1.2 + 12i
	complexValue2 := 1.2 + 12i
	complexValue3 := complex(3.2, 12)

	fmt.Println("complexValue =", complexValue)
	fmt.Println("complexValue2 =", complexValue2)
	fmt.Println("complexValue3 =", complexValue3)

	fmt.Println("complexValue3 實數 =", real(complexValue3))
	fmt.Println("complexValue3 虛數 =", imag(complexValue))
}
