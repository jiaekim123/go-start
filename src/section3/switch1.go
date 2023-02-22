package main

import "fmt"

func main() {

	// switch문 (1)
	// 제어문(조건문) - switch
	// switch 뒤 표현식(expression) 생략 가능
	// case 뒤 표현식(expression) 사용 가능
	// 자동 break 때문에 fallthrouth 존재
	// type 분기 -> 값이 아닌 변수 type으로 분기 가능

	// 예제 1
	// 개발자의 성향이지만 범위는 if로 특정 값에 떨어지는 케이스는 switch로 하는 케이스가 많음. (강사분 성향)
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 예제2
	switch b := 9; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// 예제3
	switch c := "go"; c {
	case "go":
		fmt.Println("GO!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음.")
	}

	// 예제4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음.")
	}

	// 예제5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i == j")
	case i > j:
		fmt.Println("i > j")
	}
}
