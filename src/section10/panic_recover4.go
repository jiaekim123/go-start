package main

import (
	"fmt"
	_ "log"
	"os"
)

// Panic & Recover(4)

func main() {
	// Go panic & recover 최종 실습
	fileOpen("undefined.txt")
	fmt.Println("End Main")
}

func fileOpen(filename string) {
	// defer 함수 (Panic 호출 시 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}
}
