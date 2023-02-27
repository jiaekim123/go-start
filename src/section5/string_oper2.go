package main

import "fmt"

// 데이터 타입 : 문자열 연산(2)

func main() {

	// 예제2 (비교)
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex1 : ", str1 == str2) // false
	fmt.Println("ex1 : ", str1 != str2) // true
	fmt.Println("ex1 : ", str1 > str2)  // true? -> 아님! false임!
	fmt.Println("ex1 : ", str1 < str2)  // false? -> 아님! true임!
	// Go 문자열 -> 아스키 코드에 대한 사전식 비교
	// 소문자 abcd보다 대문자 ABC가 더 크다.

}
