package main

import (
	"fmt"
	"reflect"
)

// 구조체 기본(6)

type Bicycle struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세" // 구조체 안에 구조체
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 필드 태그 사용

	// 예제1
	tag := reflect.TypeOf(Bicycle{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

	bicycle := Bicycle{
		"520d",
		"silver",
		"따릉이",
		spec{4000, 1000, 2000},
	}

	fmt.Println("ex1 : ", bicycle.name)
	fmt.Println("ex1 : ", bicycle.color)
	fmt.Println("ex1 : ", bicycle.company)
	fmt.Printf("ex1 : %#v\n", bicycle)

	fmt.Println("ex2 : ", bicycle.detail.length)
	fmt.Println("ex2 : ", bicycle.detail.height)
	fmt.Println("ex2 : ", bicycle.detail.width)
}
