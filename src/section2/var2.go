package main

import "fmt"

func main() {
	// 여러 개 묶어서 선언
	var (
		name      string = "nova"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 153
	weight = 50
	isRunning = true

	fmt.Println("name: ", name, "height : ", height, "weight : ", weight, "isRunning : ", isRunning)

}
