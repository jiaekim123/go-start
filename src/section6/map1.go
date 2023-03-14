package main

import "fmt"

// 자료형 : 맵(1)

func main() {
	// 맵(Map)
	// 맵: 해시테이블, 딕셔너리(파이썬), Key-Value 자료 저장
	// 레퍼런스 타입 (참조 값 전달)
	// 비교 연산자 사용 불가능(참조 타입 이므로)
	// 특징: 참조 타입(key)로 사용 불가능, 값(value)로는 모든 타입 사용 가능
	// make 함수 및 축약(리터럴) 초기화 가능
	// 순서 없음.

	// 예제1
	var map1 map[string]int = make(map[string]int) // 정석이지만 길어서 잘 안씀.
	var map2 = make(map[string]int)                // make로 축약해서 사용
	map3 := make(map[string]int)                   // 이 방식으로 가장 많이 사용함.

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	// 예제2
	// 방법1 한땀한땀 넣기
	map4 := map[string]int{}
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	// 방법2 json형식으로 저장
	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 33,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["orange"])
	fmt.Println("ex2 : ", map6["apple"])
}
