// 사용자 패키지 작성 및 문서화 예제
package main

import (
	"fmt"
	oper "section12/arithmetic" // alias 사용 (패키지 중복 또는 약자로 사용)
	// "section12/arithmetic"	// alias 미사용
)

func main() {
	// 사용자 패키지 작성 & GO 문서화
	// 기준은 GOPATH/src
	// 패키지 폴더명 (디렉터리 명) - 명확하게 지정
	// 패키지 파일의 Package 이름으로 사용한다. -> 길면 alias 사용
	// package main을 제외하고 package 문서에 등록
	// 지금까지 우리는 패키지를 사용해 왔다.
	// 기본적으로 GOROOT의 패키지(pkg) 검색 -> GOPATH의 패키지(src/pkg) 검색
	// go install 명령어 실행의 경우 -> GOPATH/pkg에 등록 (문서화)
	// godoc -http=:7070(임의의포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인(패키지 타입, 주석)

	// 패키지 사용 예제(사칙연산)
	nums := oper.Numbers{X: 10, Y: 2}
	fmt.Printf("ex1 : (x : %d, y : %d)\n", nums.X, nums.Y)
	fmt.Printf("ex1 : %d + %d = %d\n", nums.X, nums.Y, nums.Plus())
	fmt.Printf("ex1 : %d - %d = %d\n", nums.X, nums.Y, nums.Minus())
	fmt.Printf("ex1 : %d * %d = %d\n", nums.X, nums.Y, nums.Multi())
	fmt.Printf("ex1 : %d / %d = %d\n", nums.X, nums.Y, nums.Divide())
	fmt.Printf("ex1 : %d^ + %d^ = %d\n", nums.X, nums.Y, nums.SquarePlus())
	fmt.Printf("ex1 : %d^ - %d^ = %d\n", nums.X, nums.Y, nums.SquareMinus())
}
