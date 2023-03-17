package main

import (
	"fmt"
	"reflect"
)

// 구조체 기본(5)

type Bike struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func main() {
	// 필드 태그 사용

	// 예제1
	tag := reflect.TypeOf(Bike{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
