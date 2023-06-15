package main

import (
	"fmt"
	"reflect"
)

// 인터페이스 고급(2)
// type assertion (타입 단언/변환)

func main() {
	// 타입 변환
	// Type Assertion
	// 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// 인터페이스 (타입) -> 형 변환
	// interfaceVal.(type)

	// 예제1
	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64) // 최초의 데이터 타입으로만 갈 수 있음.

	fmt.Printf("Type: (%T)", a)
	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))

	fmt.Printf("Type: (%T)", b)
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))

	fmt.Printf("Type: (%T)", c)
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	//fmt.Printf("Type: (%T)", d)
	//fmt.Println("ex1 : ", d)

	// 예제2 (저장된 실제 타입 검사)
	if v, ok := a.(int); ok {
		fmt.Println("ex2 (int): ", v, ok)
	}
	if v, ok := a.(float64); !ok {
		fmt.Println("ex2 (float) : ", v, ok)
	}

}
