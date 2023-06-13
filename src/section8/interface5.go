package main

import "fmt"

type Teacher struct {
	name string
	age  int
}
type Students struct {
	name string
	age  int
}

// 인터페이스는 모든 값을 다 받을 수 있음.
func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {
	// 인터페이스 활용(빈 인터페이스)
	// 함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. (만능) -> 모든 타입 지정 가능

	teacher1 := Teacher{"may", 33}
	student1 := Students{"tin", 17}
	printValue(teacher1)
	printValue(student1)
	printValue(11)
	printValue(25.5)
	printValue("hello")
	printValue([]Teacher{})
	printValue([6]Students{})

}
