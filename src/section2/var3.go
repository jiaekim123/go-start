package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용 가능 (전역 x)

	// 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있음.
	shortVar1 := 100
	shortVar2 := "test"
	shortVar3 := false

	//shortVar2 := "mm" // 예외 발생

	fmt.Println("shortVar1 : ", shortVar1)
	fmt.Println("shortVar2 : ", shortVar2)
	fmt.Println("shortVar3 : ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success")
	}
}
