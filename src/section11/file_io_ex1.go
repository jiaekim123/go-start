package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 파일 I/O (1)
func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}
func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// 아래 메소드 확인 가능
	// WriteFile, ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil

	s := "hello golang!\n File Write Test!\n"

	// 파일 모드(chmod, chown, chgrp) -> 권한
	// 읽기 : 4, 쓰기 : 2, 실행 : 1
	// 소유자, 그룹, 기타 사용자 순서 (644)
	// https://golang.org/pkg/os/#FileMode

	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	check(err)

	data, err := ioutil.ReadFile("src/section11/sample.txt")
	check(err)

	fmt.Println("=========================")
	fmt.Println(string(data))
	fmt.Println("=========================")
}
