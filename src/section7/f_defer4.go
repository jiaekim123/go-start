package main

import "fmt"

func main() {
	a()
}

func a() {
	defer end(start("b")) // 중첩 함수 주의!
	fmt.Println("in a")
}

func start(message string) string {
	fmt.Println("start : ", message)
	return message
}

func end(message string) {
	fmt.Println("end: ", message)
}
