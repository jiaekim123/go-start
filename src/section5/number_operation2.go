package main

// 데이터 타입: Numeric 연산(2)

import (
	"fmt"
)

func main() {
	// 예제1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2) // 11250인데 uint8이라 242로 나옴.
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2)
	fmt.Println("ex1 : ", n1>>2)
	fmt.Println("ex1 : ", ^n1)

	// 예제2
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	//fmt.Println("ex2 : ", n3 * n4) // 에러
	fmt.Println("ex2 : ", float32(n3)+n4) // 20.2
	fmt.Println("ex2 : ", n3+int(n4))     // 20 -> 뒷자리가 잘림.
	fmt.Println("ex2 : ", n5+uint16(n6))  // 121024가 안나오고 55488이 나옴. (uint16 범위를 벗어남)

}
