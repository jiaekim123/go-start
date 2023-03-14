package main

import "fmt"

// 자료형 : 맵(2)
func main() {
	// 맵(Map)
	// 맵 조회 및 순회(Iterator)

	// 예제1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println("ex1 : ", map1["google"])
	fmt.Println()

	// 예제2(순서가 없으므로 랜덤)
	for k, v := range map1 {
		fmt.Println("ex2 : ", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex2 : ", v)
	}
}
