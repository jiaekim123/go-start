package main

func main() {
	// 기본 초기화
	// 정수 타입: 0, 실수(소수점): 0.0, 문자열: "", Boolean: True, False
	// 변수명: 숫자 첫글자 x, 대소문자 구분 O, 문자 숫자 밑줄 특수기호 사용 가능
	// 변수 및 상수 : 함수 내외 사용 가능
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi Golang!"
	var k = 4.74 // 선언 동시 초기화
	var l = "Hi! Seoul!"
	var m = true

	a = 4
	b = "hello"
	e = 77
	// 변수를 선언했는데 사용하지 않으면 에러가 남.
	println("a : ", a)
	println("b : ", b)
	println("c : ", c)
	println("d : ", d)
	println("e : ", e)
	println("f : ", f)
	println("g : ", g)
	println("h : ", h)
	println("i : ", i)
	println("j : ", j)
	println("k : ", k)
	println("l : ", l)
	println("m : ", m)

}
