package main

import "fmt"

func main() {
	// 예제1
	sayHello("Golang!")
}

func sayHello(message string) {
	defer func() {
		fmt.Println(message)
	}()

	func() {
		fmt.Println("Hi!")
	}()
}
