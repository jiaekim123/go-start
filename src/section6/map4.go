package main

import "fmt"

// 자료형 : 맵(4)

func main() {
	// 맵(Map)
	// 맵 조회 및 경우에 주의할 점

	// 예제1
	map1 := map[string]int{ // 기본 초기화 값 : int : 0, string : "", float : 0.0
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]
	value4, ok2 := map1["lemon"]

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2)     // 기본 값인 0이 나옴.
	fmt.Println("ex1 : ", value3, ok) // 두번 째 값은 키가 존재하는지 아닌지 확인 가능.
	fmt.Println("ex1 : ", value4, ok2)

	// 예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist!")
	}

	// 예제2
	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exist!")
	}

	if _, ok := map1["wiki"]; !ok {
		fmt.Println("ex2 : kiwi is not true")
	}
}
