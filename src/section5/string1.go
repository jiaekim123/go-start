package main

// 데이터 타입 : 문자열(1)

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// 큰 따옴표 "", 백스쿼트 ₩₩
	// golang: 문자 char 타입 제공 안함. rune (int32)로 문자 코드 값으로
	// 문자: '' 로 작성
	// 자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)...

	// 예제1
	var str1 string = "c:\\go-start\\src\\" // c:\go_study\src
	str2 := `c:\go_study\src\`              // 가독성을 위해 문자열의 백스쿼트를 사용해서 보이는 그대로를 표현할 수도 있음.`

	fmt.Println("ex1 : ", str1)
	fmt.Println("ex1 : ", str2)

	// 예제2
	var str3 string = "Hello world!"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00"

	fmt.Println("ex2 : ", str3)
	fmt.Println("ex2 : ", str4)
	fmt.Println("ex2 : ", str5)

	// 예제3
	// 길이 (바이트 수) -> 영어: 1바이트, 한글: 3바이트
	fmt.Println("ex3 : ", len(str3))
	fmt.Println("ex3 : ", len(str4))

	// 예제4
	// 길이 (실제 길이)
	fmt.Println("ex4 : ", utf8.RuneCountInString(str3))
	fmt.Println("ex4 : ", utf8.RuneCountInString(str4)) // utf8 패키지를 통해 구하는 것을 선호함.
	fmt.Println("ex4 : ", len([]rune(str4)))            // len으로 실제 길이 구하는 법
}
