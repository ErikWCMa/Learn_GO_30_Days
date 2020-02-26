package main

import "fmt"

func changeTest(test string, x string) string {
	test = x
	fmt.Println("test in change ", test)
	return test
}

func main() {
	var test string = "gnip"

	message := make(chan string)
	go func() { message <- "333" }()
	go func() { message <- test }()
	go changeTest(test, "ping")
	msg := <-message
	fmt.Println("msg:", msg)
	fmt.Println(test)
	test = changeTest(test, "ping")
	fmt.Println(test)
}
