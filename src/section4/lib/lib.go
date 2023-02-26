package lib

import "fmt"

// 라이브러리 접근 제어
// 패키지 명은 폴더 이름

func CheckNum(c int32) bool {
	return c > 10
}

func init() {
	fmt.Println("lib Pakage > init start!")
}
