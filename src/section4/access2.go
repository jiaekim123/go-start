package main

import (
	"fmt"
	checkUp1 "section4/lib"
	_ "section4/lib2" // 지금은 안쓰지만 컴파일러에서 삭제하지 않고 패스
)

// 패키지 접근 제어2
func main() {
	// 패키지 접근 제어
	// 별칭 사용
	// 빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUp1.CheckNum(1010))
}
