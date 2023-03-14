package main

import "fmt"

// 자료형 : 맵(3)
func main() {
	// 맵(Map)
	// 맵 값 변경 및 삭제

	// 예제1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex1 : ", map1)

	map1["home2"] = "http://test1.com" // 추가

	fmt.Println("ex1 : ", map1)

	map1["home2"] = "http://test2.com" // 변경

	fmt.Println("ex1 : ", map1)

	// 예제2 (삭제)
	delete(map1, "home2")
	fmt.Println("ex1 : ", map1)

}
