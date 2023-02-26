package main

import (
	"fmt"
	_ "section4/lib"
)

func init() {
	fmt.Println("init1 method start!")
}

func init() {
	fmt.Println("init2 method start!")
}

func init() {
	fmt.Println("init3 method start!")
}

func main() {

	fmt.Println("main method start!")
}
