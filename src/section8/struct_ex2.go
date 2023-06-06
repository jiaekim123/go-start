package main

import "fmt"

type Account2 struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account2) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account2) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// 예제1
	kim := Account2{"245-901", 10000, 0.015}
	lee := Account2{"245-902", 20000, 0.025}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)
	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))
}
