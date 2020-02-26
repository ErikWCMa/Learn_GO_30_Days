package main

import (
	"fmt"
	"time"
	//"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func e(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	go e(2)
	go fmt.Println("123")
	time.Sleep(time.Millisecond * 1)
}

// func main() {
// 	f(0)
// 	e(2)
// }
// 有順序的執行
